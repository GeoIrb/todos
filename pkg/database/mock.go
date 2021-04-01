package database

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/geoirb/todos/pkg/storage"
)

// DatabaseMock ...
type DatabaseMock struct {
	mock.Mock
}

var _ User = &DatabaseMock{}

// Insert ...
func (m *DatabaseMock) Insert(ctx context.Context, user storage.UserInfo) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// SelectOne ...
func (m *DatabaseMock) SelectOne(ctx context.Context, filter storage.UserFilter) (storage.UserInfo, error) {
	args := m.Called(ctx, filter)
	if user, ok := args.Get(0).(storage.UserInfo); ok {
		return user, args.Error(0)
	}
	return storage.UserInfo{}, nil
}

// SelectList ...
func (m *DatabaseMock) SelectList(ctx context.Context, filter storage.UserFilter) ([]storage.UserInfo, error) {
	args := m.Called(ctx, filter)
	if users, ok := args.Get(0).([]storage.UserInfo); ok {
		return users, args.Error(0)
	}
	return nil, nil
}
