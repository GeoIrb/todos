package postgresql

import (
	"context"
	"fmt"
	"sync"

	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Task database.
type Task struct {
	mutex sync.Mutex
	db    *sqlx.DB

	insertTask    string
	selectTask    string
	selectOrderBy string
	updateTask    string
	deleteTask    string

	connect func() (*sqlx.DB, error)
}

var _ database.Task = &Task{}

func NewTask(
	connectLayout string,
	host string,
	port int,
	database string,
	user string,
	password string,

	insertTask string,
	selectTask string,
	selectOrderBy string,
	updateTask string,
	deleteTask string,
) (t *Task, err error) {
	t = &Task{
		insertTask:    insertTask,
		selectTask:    selectTask,
		selectOrderBy: selectOrderBy,
		updateTask:    updateTask,
		deleteTask:    deleteTask,
	}
	connectCfg := fmt.Sprintf(connectLayout, host, port, user, password, database)
	t.connect = func() (*sqlx.DB, error) {
		return sqlx.Connect("postgres", connectCfg)
	}
	t.db, err = t.connect()
	return
}

func (t *Task) Insert(ctx context.Context, task storage.TaskInfo) (err error) {
	if err = t.check(); err != nil {
		return
	}

	_, err = t.db.QueryContext(ctx, t.insertTask, task.UserID, task.Title, task.Description, task.Deadline)
	return
}
func (t *Task) Select(ctx context.Context, filter storage.TaskFilter) (tasks []storage.TaskInfo, err error) {
	if err = t.check(); err != nil {
		return
	}
	query := t.selectTask

	if filter.ID != nil {
		query += fmt.Sprintf(" AND id = %d", *filter.ID)
	}

	if filter.From != nil && filter.To != nil {
		query += fmt.Sprintf("AND deadline BETWEEN %d AND %d", *filter.From, *filter.To)
	}

	query += fmt.Sprintf(" ORDER BY %s", t.selectOrderBy)

	var dbTasks []taskInfo
	if err = t.db.SelectContext(ctx, &dbTasks, query, filter.UserID); err == nil {
		tasks = make([]storage.TaskInfo, 0, len(dbTasks))
		for _, task := range dbTasks {
			tasks = append(tasks, storage.TaskInfo(task))
		}
	}
	return
}
func (t *Task) Update(ctx context.Context, task storage.TaskInfo) (err error) {
	if err = t.check(); err != nil {
		return
	}

	_, err = t.db.QueryContext(ctx, t.updateTask, task.Title, task.Description, task.Deadline, task.ID)
	return
}
func (t *Task) Delete(ctx context.Context, filter storage.TaskFilter) (err error) {
	if err = t.check(); err != nil {
		return
	}

	if filter.ID == nil {
		return errNotFoundParam
	}

	_, err = t.db.QueryContext(ctx, t.deleteTask, *filter.ID)
	return
}

func (t *Task) Close() error {
	return t.db.Close()
}

func (t *Task) check() (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if err = t.db.Ping(); err != nil {
		if t.db, err = t.connect(); err != nil {
			err = fmt.Errorf("connect db %s", err)
		}
	}
	return
}
