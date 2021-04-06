package todos

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) CreateTask(ctx context.Context, task TaskInfo) (err error) {
	args := m.Called(task)
	return args.Error(0)
}
func (m *ServiceMock) UpdateTask(ctx context.Context, task TaskInfo) (err error) {
	args := m.Called(task)
	return args.Error(0)
}
func (m *ServiceMock) DeleteTask(ctx context.Context, filter Filter) (err error) {
	args := m.Called(filter)
	return args.Error(0)
}
func (m *ServiceMock) GetTaskList(ctx context.Context, filter Filter) (tasks []TaskInfo, err error) {
	args := m.Called(filter)

	var ok bool
	if tasks, ok = args.Get(0).([]TaskInfo); ok {
		err = args.Error(0)
	}
	return
}
