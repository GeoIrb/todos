package sender

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// SenderMock ...
type SenderMock struct {
	mock.Mock
}

func (m *SenderMock) Send(ctx context.Context, dst, message string) (err error) {
	args := m.Called(dst, message)
	return args.Error(0)
}
