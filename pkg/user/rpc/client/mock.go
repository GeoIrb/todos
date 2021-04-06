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
func (m *AuthClientMock) Authorization(ctx context.Context, token string) (id int, err error) {
	args := m.Called(token)

	var ok bool
	if id, ok = args.Get(0).(int); ok {
		err = args.Error(1)
		return
	}
	return
}
