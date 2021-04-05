package main

import (
	"crypto/sha256"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fasthttp/router"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"

	"github.com/geoirb/todos/pkg/cache/freecache"
	"github.com/geoirb/todos/pkg/database/postgresql"
	"github.com/geoirb/todos/pkg/jwt"
	"github.com/geoirb/todos/pkg/password"
	"github.com/geoirb/todos/pkg/sender/smtp"
	userStorage "github.com/geoirb/todos/pkg/storage/user"
	"github.com/geoirb/todos/pkg/token"
	"github.com/geoirb/todos/pkg/user"
	"github.com/geoirb/todos/pkg/user/http"
	"github.com/geoirb/todos/pkg/user/rpc"
	"github.com/geoirb/todos/pkg/user/rpc/server"
)

type configuration struct {
	HttpPort string `envconfig:"HTTP_PORT" default:"8080"`
	RpcPort  string `envconfig:"RPC_PORT" default:"8070"`

	DBConnectLayout  string `envconfig:"DB_CONNECT_LAYOUT" default:"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"`
	DBHost           string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort           int    `envconfig:"DB_PORT" default:"5432"`
	DBName           string `envconfig:"DB_NAME" default:"todos"`
	DBUser           string `envconfig:"DB_USER" default:"secret-user"`
	DBPassword       string `envconfig:"DB_PASSWORD" default:"secret-password"`
	DBInsertUser     string `envconfig:"DB_INSERT_USER" default:"INSERT INTO public.user(email, password, is_active)VALUES ($1, $2, $3);"`
	DBSelectUser     string `envconfig:"DB_SELECT_USER" default:"SELECT * FROM public.user WHERE email=$1 AND password=$2"`
	DBSelectUserList string `envconfig:"DB_SELECT_USER_LIST" default:"SELECT * FROM public.user"`

	StoragePasswordTTL time.Duration `envconfig:"STORAGE_PASSWORD_TTL" default:"24h"`

	SMTPAddress        string        `envconfig:"SMTP_ADDRESS" default:"tokenplace21@yandex.ru"`
	SMTPPassword       string        `envconfig:"SMTP_PASSWORD" default:"21tokenplace21"`
	SMTPHost           string        `envconfig:"SMTP_HOST" default:"smtp.yandex.ru"`
	SMTPPort           int           `envconfig:"SMTP_PORT" default:"465"`
	SMTPConnectTimeout time.Duration `envconfig:"SMTP_CONNECT_TIMEOUT" default:"3s"`
	SMTPSendTimeout    time.Duration `envconfig:"SMTP_SEND_TIMEOUT" default:"5s"`

	PasswordSecret []byte `envconfig:"PASSWORD_SECRET" default:"password-secret"`

	JWTTokenTTL time.Duration `envconfig:"JWT_TOKEN_TTL" default:"1h"`
	JWTSecret   []byte        `envconfig:"JWT_SECRET" default:"jwt-secret"`
}

const (
	prefixCfg   = ""
	serviceName = "user-service"
)

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.WithPrefix(logger, "service", serviceName)

	var cfg configuration
	if err := envconfig.Process(prefixCfg, &cfg); err != nil {
		level.Error(logger).Log("msg", "configuration", "err", err)
		os.Exit(1)
	}

	level.Info(logger).Log("msg", "initializing")

	db, err := postgresql.NewUser(
		cfg.DBConnectLayout,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBInsertUser,
		cfg.DBSelectUser,
		cfg.DBSelectUserList,
	)
	if err != nil {
		level.Error(logger).Log("msg", "db init", "err", err)
		os.Exit(1)
	}

	cache := freecache.NewUser(123)

	storage := userStorage.NewStorage(
		db,
		cache,
		cfg.StoragePasswordTTL,
	)

	smtp := smtp.New(
		cfg.SMTPAddress,
		cfg.SMTPPassword,
		cfg.SMTPHost,
		cfg.SMTPPort,
		cfg.SMTPConnectTimeout,
		cfg.SMTPSendTimeout,
	)

	hash := password.NewHash(
		sha256.New,
		cfg.PasswordSecret,
	)

	token := token.New()

	jwt := jwt.New(
		cfg.JWTTokenTTL,
		cfg.JWTSecret,
	)

	svc := user.NewService(
		storage,
		smtp,
		hash,
		token,
		jwt,
		logger,
	)

	router := router.New()
	http.Routing(router, svc, token)

	httpServer := &fasthttp.Server{
		Handler:          router.Handler,
		DisableKeepalive: true,
	}

	go func() {
		level.Info(logger).Log("msg", "http server turn on", "port", cfg.HttpPort)
		if err := httpServer.ListenAndServe(":" + cfg.HttpPort); err != nil {
			level.Error(logger).Log("msg", "http server turn on", "err", err)
			os.Exit(1)
		}
	}()

	rpcServer := grpc.NewServer()
	defer rpcServer.Stop()
	rpc.RegisterAuthServer(
		rpcServer,
		server.NewAuthRPCServer(svc),
	)

	lis, err := net.Listen("tcp", ":"+cfg.RpcPort)
	if err != nil {
		level.Error(logger).Log("msg", "failed to turn up tcp connection", "err", err)
		os.Exit(1)
	}
	defer lis.Close()

	go func() {
		level.Info(logger).Log("msg", "rpc server turn on", "port", cfg.RpcPort)
		if err := rpcServer.Serve(lis); err != nil {
			level.Error(logger).Log("msg", "rpc auth server turn on", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	level.Info(logger).Log("msg", "received signal, exiting signal", "signal", <-c)

	if err := httpServer.Shutdown(); err != nil {
		level.Info(logger).Log("msg", "http server shoutdown", "err", err)
	}
}
