package postgresql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"

	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
)

// User database.
type User struct {
	mutex sync.Mutex
	db    *sqlx.DB

	connect func() (*sqlx.DB, error)
}

var _ database.User = &User{}

func NewUser(
	dbDriver string,
	connectLayout string,
	host string,
	port int,
	database string,
	user string,
	password string,
) (u *User, err error) {
	u = &User{}
	connectCfg := fmt.Sprintf(connectLayout, host, port, user, password, database)

	u.connect = func() (*sqlx.DB, error) {
		return sqlx.Connect(dbDriver, connectCfg)
	}
	u.db, err = u.connect()
	return
}

func (u *User) Insert(ctx context.Context, user storage.UserInfo) (err error) {
	if err = u.check(); err != nil {
		return
	}
	return
}

func (u *User) Select(ctx context.Context, filter storage.UserFilter) (users []storage.UserInfo, err error) {
	if err = u.check(); err != nil {
		return
	}
	return
}

func (u *User) Close() error {
	return u.db.Close()
}

func (u *User) check() (err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	if err = u.db.Ping(); err != nil {
		if u.db, err = u.connect(); err != nil {
			err = fmt.Errorf("connect db %s", err)
		}
	}
	return
}
