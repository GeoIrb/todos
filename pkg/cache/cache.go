package cache

import (
	"context"
	"time"
)

// User cache.
type User interface {
	SetPassword(ctx context.Context, mail, password string, ttl time.Duration) (err error)
	GetPassword(ctx context.Context, mail string) (password string, isExist bool, err error)
}
