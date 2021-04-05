package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"

	"github.com/geoirb/todos/pkg/database/postgresql"
	taskStorage "github.com/geoirb/todos/pkg/storage/task"
	"github.com/geoirb/todos/pkg/todos"
	"github.com/geoirb/todos/pkg/todos/http"
	"github.com/geoirb/todos/pkg/token"
	"github.com/geoirb/todos/pkg/user/rpc"
	"github.com/geoirb/todos/pkg/user/rpc/client"
)

type configuration struct {
	HttpPort string `envconfig:"HTTP_PORT" default:"8081"`

	AuthHost string `envconfig:"AUTH_HOST" default:"127.0.0.1"`
	AuthPort int    `envconfig:"AUTH_PORT" default:"8070"`

	DBConnectLayout string `envconfig:"DB_CONNECT_LAYOUT" default:"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"`
	DBHost          string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort          int    `envconfig:"DB_PORT" default:"5432"`
	DBName          string `envconfig:"DB_NAME" default:"todos"`
	DBUser          string `envconfig:"DB_USER" default:"secret-user"`
	DBPassword      string `envconfig:"DB_PASSWORD" default:"secret-password"`
	DBInsertTask    string `envconfig:"DB_INSERT_TASK" default:"INSERT INTO public.task(user_id, title, description, deadline)VALUES ($1, $2, $3, $4)"`
	DBSelectTask    string `envconfig:"DB_SELECT_TASK" default:"SELECT * FROM public.task WHERE user_id=$1"`
	DBSelectOrderBy string `envconfig:"DB_SELECT_ORDER_BY" default:"deadline"`
	DBUpdateTask    string `envconfig:"DB_SELECT_TASK" default:"UPDATE public.task SET title = $1, description = $2, deadline = $3 WHERE id = $4"`
	DBDeleteTask    string `envconfig:"DB_DELETE_TASK" default:"DELETE FROM public.task WHERE id = $1 AND user_id = $2"`
}

const (
	prefixCfg   = ""
	serviceName = "todos-service"
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

	conn, err := grpc.Dial(
		fmt.Sprint(cfg.AuthHost, ":", cfg.AuthPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		level.Error(logger).Log("msg", "grpc connect", "err", err)
		os.Exit(1)
	}
	defer conn.Close()

	auth := client.NewAuthRPCClient(rpc.NewAuthClient(conn))

	db, err := postgresql.NewTask(
		cfg.DBConnectLayout,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBInsertTask,
		cfg.DBSelectTask,
		cfg.DBSelectOrderBy,
		cfg.DBUpdateTask,
		cfg.DBDeleteTask,
	)
	if err != nil {
		level.Error(logger).Log("msg", "db init", "err", err)
		os.Exit(1)
	}

	storage := taskStorage.NewStorage(
		db,
	)

	token := token.New()

	svc := todos.NewService(
		auth,
		storage,
		token,

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	level.Info(logger).Log("msg", "received signal, exiting signal", "signal", <-c)

	if err := httpServer.Shutdown(); err != nil {
		level.Info(logger).Log("msg", "http server shoutdown", "err", err)
	}
}
