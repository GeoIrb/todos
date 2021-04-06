package todos

import (
	"context"
)

// Service todos.
type Service interface {
	CreateTask(ctx context.Context, task TaskInfo) (err error)
	UpdateTask(ctx context.Context, task TaskInfo) (err error)
	DeleteTask(ctx context.Context, filter Filter) (err error)
	GetTaskList(ctx context.Context, filter Filter) (tasks []TaskInfo, err error)
}
