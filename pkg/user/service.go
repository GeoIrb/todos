package user

import (
	"context"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sethvargo/go-password/password"

	"github.com/geoirb/todos/pkg/sender"
	"github.com/geoirb/todos/pkg/storage"
)

type token interface {
	Get(ctx context.Context) (token string)
}

type hash interface {
	Password(ctx context.Context, password string) (passwordHash string)
}

type jwt interface {
	CreateToken(ctx context.Context, id string) (token string, err error)
	Parse(ctx context.Context, token string) (isValid bool, id string, err error)
}

type Service struct {
	storage storage.User
	email   sender.Email

	hash  hash
	token token
	jwt   jwt

	logger log.Logger
}

// Registration new user in system.
func (s *Service) Registration(ctx context.Context, info Registartion) error {
	logger := log.WithPrefix(s.logger, "method", "Registration", "email", info.Email)

	filter := storage.UserFilter{
		Email: &info.Email,
	}
	users, err := s.storage.Select(ctx, filter)
	if err != nil {
		level.Error(logger).Log("msg", "storage select", "err", err)
		return err
	}
	if len(users) != 0 {
		level.Error(logger).Log("err", "user is exist")
		return ErrUserIsExist
	}

	user := storage.UserInfo{
		Email: info.Email,
	}

	if user.Password, err = password.Generate(8, 3, 0, false, false); err != nil {
		level.Error(logger).Log("msg", "password generate", "err", err)
		return err
	}

	if err = s.email.Send(ctx, user.Email, "Temporary password "+user.Password); err != nil {
		level.Error(logger).Log("msg", "send password", "err", err)
		return err
	}

	return s.storage.New(ctx, user)
}

// Login user in system.
func (s *Service) Login(ctx context.Context, info Login) (auth Auth, err error) {
	logger := log.WithPrefix(s.logger, "method", "Login", "email", info.Email)

	hashPassword := s.hash.Password(ctx, info.Password)
	filter := storage.UserFilter{
		Email:    &info.Email,
		Password: &hashPassword,
	}

	user, err := s.storage.Get(ctx, filter)
	if err != nil {
		level.Error(logger).Log("msg", "storage get", "err", err)
		return
	}

	if !user.IsActive {
		level.Info(logger).Log("msg", "user is not active")
		if user.Password != info.Password {
			err = ErrUserNotFound
		}
		return
	}

	token, err := s.jwt.CreateToken(ctx, user.ID)
	auth.Token = &token
	return
}

// Create user in system.
func (s *Service) Create(ctx context.Context, info Create) (err error) {
	logger := log.WithPrefix(s.logger, "method", "Create", "email", info.Email)

	filter := storage.UserFilter{
		Email: &info.Email,
	}

	user, err := s.storage.Get(ctx, filter)
	if err != nil {
		level.Error(logger).Log("msg", "storage get", "err", err)
		return
	}

	if user.Password != info.OldPassword {
		err = ErrUserNotFound
		level.Error(logger).Log("err", err)
		return
	}

	user.Password = s.hash.Password(ctx, info.NewPassword)
	return s.storage.Create(ctx, user)
}

// Authorization token.
func (s *Service) Authorization(ctx context.Context, token string) (id string, err error) {
	parts := strings.Split(token, " ")
	if len(parts) != 2 && parts[0] != "Bearer" {
		err = ErrFiledAuthenticate
		return
	}

	var isValid bool
	if isValid, id, err = s.jwt.Parse(ctx, parts[1]); !isValid && err != nil {
		err = ErrTokenExpired
		return
	}
	return
}

// GetUserList by filter.
func (s *Service) GetUserList(ctx context.Context, filter Filter) (users []User, err error) {
	logger := log.WithPrefix(s.logger, "method", "GetUserList")

	token := s.token.Get(ctx)
	if _, err = s.Authorization(ctx, token); err != nil {
		level.Error(logger).Log("msg", "authorization", "err", err)
		return
	}

	userFilter := storage.UserFilter{
		ID:    filter.ID,
		Email: filter.Email,
	}

	storageUsers, err := s.storage.Select(ctx, userFilter)
	if err != nil {
		level.Error(logger).Log("msg", "storage select", "err", err)
		return
	}

	users = make([]User, 0, len(storageUsers))
	for _, user := range storageUsers {
		users = append(
			users,
			User{
				ID:    user.ID,
				Email: user.Email,
			},
		)
	}
	return
}
