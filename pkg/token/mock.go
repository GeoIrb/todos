package token

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// TokenMock ...
type TokenMock struct {
	mock.Mock
}

func (m *TokenMock) Put(ctx context.Context, token string) context.Context {
	args := m.Called(token)
	ctx, _ = args.Get(0).(context.Context)
	return ctx
}

func (m *TokenMock) Get(ctx context.Context) string {
	args := m.Called()
	str, _ := args.Get(0).(string)
	return str
}
