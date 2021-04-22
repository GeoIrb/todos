package todos

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/geoirb/todos/pkg/storage"
	"github.com/geoirb/todos/pkg/user/rpc/client"
)

type token interface {
	Get(ctx context.Context) (token string)
}

type service struct {
	auth    client.Auth
	storage storage.Task
	token   token

	logger log.Logger
}

func NewService(
	auth client.Auth,
	storage storage.Task,
	token token,

	logger log.Logger,
) Service {
	return &service{
		auth:    auth,
		storage: storage,
		token:   token,

		logger: logger,
	}
}

func (s *service) CreateTask(ctx context.Context, task TaskInfo) (err error) {
	logger := log.WithPrefix(s.logger, "method", "NewTask")

	token := s.token.Get(ctx)
	if task.UserID, err = s.auth.Authorization(ctx, token); err != nil {
		return ErrFailedAuthenticate
	}

	if err = s.storage.Create(ctx, storage.TaskInfo(task)); err != nil {
		level.Error(logger).Log("msg", "storage new task", "err", err)
	}
	return
}

func (s *service) UpdateTask(ctx context.Context, task TaskInfo) (err error) {
	logger := log.WithPrefix(s.logger, "method", "UpdateTask")

	token := s.token.Get(ctx)
	if task.UserID, err = s.auth.Authorization(ctx, token); err != nil {
		return ErrFailedAuthenticate
	}

	if err = s.storage.Update(ctx, storage.TaskInfo(task)); err != nil {
		level.Error(logger).Log("msg", "storage new task", "err", err)
	}
	return
}

func (s *service) DeleteTask(ctx context.Context, filter Filter) (err error) {
	logger := log.WithPrefix(s.logger, "method", "DeleteTask")

	token := s.token.Get(ctx)
	if _, err = s.auth.Authorization(ctx, token); err != nil {
		return ErrFailedAuthenticate
	}

	if err = s.storage.Delete(ctx, storage.TaskFilter(filter)); err != nil {
		level.Error(logger).Log("msg", "storage new task", "err", err)
	}
	return
}

func (s *service) GetTaskList(ctx context.Context, filter Filter) (tasks []TaskInfo, err error) {
	logger := log.WithPrefix(s.logger, "method", "GetTaskList")
	token := s.token.Get(ctx)
	if filter.UserID, err = s.auth.Authorization(ctx, token); err != nil {
		err = ErrFailedAuthenticate
		return
	}

	storageTasks, err := s.storage.GetList(ctx, storage.TaskFilter(filter))
	if err != nil {
		level.Error(logger).Log("msg", "storage new task", "err", err)
		return
	}
	tasks = make([]TaskInfo, 0, len(storageTasks))
	for _, task := range storageTasks {
		tasks = append(tasks, TaskInfo(task))
	}
	return
}
