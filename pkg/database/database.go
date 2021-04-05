package database

import (
	"context"

	"github.com/geoirb/todos/pkg/storage"
)

// User database.
type User interface {
	Insert(ctx context.Context, user storage.UserInfo) error
	SelectOne(ctx context.Context, filter storage.UserFilter) (storage.UserInfo, error)
	SelectList(ctx context.Context, filter storage.UserFilter) ([]storage.UserInfo, error)
}

// Task database.
type Task interface {
	Insert(ctx context.Context, task storage.TaskInfo) error
	Select(ctx context.Context, filter storage.TaskFilter) ([]storage.TaskInfo, error)
	Update(ctx context.Context, task storage.TaskInfo) error
	Delete(ctx context.Context, filter storage.TaskFilter) error
}
