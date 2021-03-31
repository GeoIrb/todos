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

// Create saves user info..
func (s *Storage) Create(ctx context.Context, user storage.UserInfo) error {
	return s.db.Insert(ctx, user)
}

func (s *Storage) Get(ctx context.Context, filter storage.UserFilter) ([]storage.UserInfo, error) {
	if password, isExist, err := s.cache.GetPassword(ctx, mail); isExist || err != nil{
		
	}
	return s.db.Select(ctx, filter)
}
