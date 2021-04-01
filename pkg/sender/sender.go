package sender

import (
	"context"
)

// Email sender
type Email interface {
	Send(ctx context.Context, dst, message string) (err error)
}
