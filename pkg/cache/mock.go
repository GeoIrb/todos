package cache

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// CacheMock ...
type CacheMock struct {
	mock.Mock
}

// SetPassword ...
func (m *CacheMock) SetPassword(ctx context.Context, email, password string, ttl time.Duration) (err error) {
	args := m.Called(email, password, ttl)
	return args.Error(0)
}

// GetPassword ...
func (m *CacheMock) GetPassword(ctx context.Context, email string) (password string, isExist bool, err error) {
	args := m.Called(email)

	var ok bool
	if password, ok = args.Get(0).(string); !ok {
		return
	}

	if isExist, ok = args.Get(1).(bool); !ok {
		return
	}

	err = args.Error(0)
	return
}

func (m *CacheMock) DeletePassword(ctx context.Context, email string) {
	m.Called(email)
}
