package user

import (
	"context"
	"time"

	"github.com/geoirb/todos/pkg/cache"
	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
)

// Storage for user.
type Storage struct {
	db    database.User
	cache cache.User

	passwordTTL time.Duration
}

var _ storage.User = &Storage{}

// New temporarily saves the password for the new user.
func (s *Storage) New(ctx context.Context, user storage.UserInfo) error {
	return s.cache.SetPassword(ctx, user.Email, user.Password, s.passwordTTL)
}

// Create saves user info..
func (s *Storage) Create(ctx context.Context, user storage.UserInfo) (storage.UserInfo, error) {
	u, err := s.db.Insert(ctx, database.UserInfo(user))
	return storage.UserInfo(u), err
}

func (s *Storage) Get(ctx context.Context, filter storage.UserFilter) (storage.UserInfo, error) {
	return storage.UserInfo{}, nil
}

func NewStorage(
	db database.User,
	cache cache.User,

	passwordTTL time.Duration,
) *Storage {
	return &Storage{
		db:    db,
		cache: cache,

		passwordTTL: passwordTTL,
	}
}
