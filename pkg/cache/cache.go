package cache

import (
	"context"
	"time"
)

// User cache.
type User interface {
	SetPassword(ctx context.Context, email, password string, ttl time.Duration) (err error)
	GetPassword(ctx context.Context, email string) (password string, isExist bool, err error)
	DeletePassword(ctx context.Context, email string)
}
