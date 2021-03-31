package cache

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type CacheMock struct {
	mock.Mock
}

func (m *CacheMock) SetPassword(ctx context.Context, mail, password string, ttl time.Duration) (err error) {
	args := m.Called(ctx, mail, password, ttl)
	return args.Error(0)
}
func (m *CacheMock) GetPassword(ctx context.Context, mail string) (password string, isExist bool, err error) {
	args := m.Called(ctx, mail)

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
