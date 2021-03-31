package user

import (
	"context"

	"github.com/geoirb/todos/pkg/cache"
	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
)

// Storage for user.
type Storage struct {
	db    database.User
	cache cache.User
}

var _ storage.User = &Storage{}

func (s *Storage) New(ctx context.Context, user storage.UserInfo) error { return nil }
func (s *Storage) Create(ctx context.Context, user storage.UserInfo) (storage.UserInfo, error) {
	return storage.UserInfo{}, nil
}
func (s *Storage) Get(ctx context.Context, filter storage.UserFilter) (storage.UserInfo, error) {
	return storage.UserInfo{}, nil
}

func NewStorage() {}
