package storage

import (
	"context"
)

// User storage.
type User interface {
	New(ctx context.Context, user UserInfo) error
	Create(ctx context.Context, user UserInfo) (UserInfo, error)
	Get(ctx context.Context, filter UserFilter) (UserInfo, error)
}
