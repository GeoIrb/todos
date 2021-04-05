package task

import (
	"context"

	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
)

// Storage task.
type Storage struct {
	db database.Task
}

var _ storage.Task = &Storage{}

// NewStorage return storage for task.
func NewStorage(
	db database.Task,
) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Create(ctx context.Context, task storage.TaskInfo) error {
	return s.db.Insert(ctx, task)
}

func (s *Storage) GetList(ctx context.Context, filter storage.TaskFilter) ([]storage.TaskInfo, error) {
	return s.db.Select(ctx, filter)
}

func (s *Storage) Update(ctx context.Context, task storage.TaskInfo) error {
	return s.db.Update(ctx, task)
}

func (s *Storage) Delete(ctx context.Context, filter storage.TaskFilter) error {
	return s.db.Delete(ctx, filter)
}
