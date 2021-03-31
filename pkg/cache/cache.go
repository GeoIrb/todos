package cache

import (
	"context"
	"time"
)

// User cache.
type User interface {
	SetPassword(ctx context.Context, key, password string, ttl time.Duration) (err error)
	GetPassword(ctx context.Context, key string) (password string, isExist bool, err error)
}
