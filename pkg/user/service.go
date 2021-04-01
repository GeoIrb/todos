package user

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sethvargo/go-password/password"

	"github.com/geoirb/todos/pkg/sender"
	"github.com/geoirb/todos/pkg/storage"
)

type hash interface {
	Password(ctx context.Context, password string) (passwordHash string)
}

type jwt interface {
	CreateToken(ctx context.Context, id string) (token string, err error)
	Parse(ctx context.Context, token string) (isValid bool, id string, err error)
}

type Service struct {
	storage storage.User
	hash    hash
	jwt     jwt
	email   sender.Email

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

func (s *Service) Login(ctx context.Context, info Login) (Auth, error) {
	logger := log.WithPrefix(s.logger, "method", "Login", "email", info.Email)

	hashPassword := s.hash.Password(ctx, info.Password)
	filter := storage.UserFilter{
		Email: &info.Email,
		Password: &hashPassword,
	}

	s.storage.Get(ctx, filter)
}
