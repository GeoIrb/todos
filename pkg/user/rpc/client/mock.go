package client

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// AuthClientMock ...
type AuthClientMock struct {
	mock.Mock
}

var _ Auth = &AuthClientMock{}

// Authorization ...
func (m *AuthClientMock) Authorization(ctx context.Context, token string) (id string, err error) {
	args := m.Called(ctx, token)

	var ok bool
	if id, ok = args.Get(0).(string); ok {
		err = args.Error(1)
		return
	}
	return
}
