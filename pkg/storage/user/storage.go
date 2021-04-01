package user

import (
	"context"
	"fmt"
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

// NewStorage return storage for user.
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

// New temporarily saves the password for the new user.
func (s *Storage) New(ctx context.Context, user storage.UserInfo) error {
	return s.cache.SetPassword(ctx, user.Email, user.Password, s.passwordTTL)
}

// Create saves user info.
func (s *Storage) Create(ctx context.Context, user storage.UserInfo) error {
	if err := s.db.Insert(ctx, user); err != nil {
		return err
	}
	s.cache.DeletePassword(ctx, user.Email)
	return nil
}

// Get user by filter.
func (s *Storage) Get(ctx context.Context, filter storage.UserFilter) (user storage.UserInfo, err error) {
	if filter.Email == nil || filter.Password == nil {
		err = fmt.Errorf("not found params")
		return
	}
	password, isExist, err := s.cache.GetPassword(ctx, *filter.Email)
	if isExist {
		user.Email = *filter.Email
		user.Password = password
		return user, err
	}
	return s.db.SelectOne(ctx, filter)
}

// Select users by filter.
func (s *Storage) Select(ctx context.Context, filter storage.UserFilter) ([]storage.UserInfo, error) {
	return s.db.SelectList(ctx, filter)
}
