package storage

import (
	"context"
)

// User storage.
type User interface {
	New(ctx context.Context, user UserInfo) error
	Create(ctx context.Context, user UserInfo) error
	Get(ctx context.Context, filter UserFilter) (UserInfo, error)
	Select(ctx context.Context, filter UserFilter) ([]UserInfo, error)
}

type Task interface {
	Create(ctx context.Context, task TaskInfo) error
	GetList(ctx context.Context, filter TaskFilter) ([]TaskInfo, error)
	Update(ctx context.Context, task TaskInfo) error
	Delete(ctx context.Context, filter TaskFilter) error
}
