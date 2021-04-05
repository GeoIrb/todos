package postgresql

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/geoirb/todos/pkg/database"
	"github.com/geoirb/todos/pkg/storage"
)

// User database.
type User struct {
	mutex sync.Mutex
	db    *sqlx.DB

	insertUser     string
	selectUser     string
	selectUserList string

	connect func() (*sqlx.DB, error)
}

var _ database.User = &User{}

func NewUser(
	connectLayout string,
	host string,
	port int,
	database string,
	user string,
	password string,

	insertUser string,
	selectUser string,
	selectUserList string,
) (u *User, err error) {
	u = &User{
		insertUser:     insertUser,
		selectUser:     selectUser,
		selectUserList: selectUserList,
	}
	connectCfg := fmt.Sprintf(connectLayout, host, port, user, password, database)
	u.connect = func() (*sqlx.DB, error) {
		return sqlx.Connect("postgres", connectCfg)
	}
	u.db, err = u.connect()
	return
}

func (u *User) Insert(ctx context.Context, user storage.UserInfo) (err error) {
	if err = u.check(); err != nil {
		return
	}
	_, err = u.db.QueryContext(ctx, u.insertUser, user.Email, user.Password, true)
	return
}

func (u *User) SelectOne(ctx context.Context, filter storage.UserFilter) (user storage.UserInfo, err error) {
	if err = u.check(); err != nil {
		return
	}

	if filter.Email == nil || filter.Password == nil {
		err = errNotFoundParam
		return
	}

	var dbUser userInfo
	err = u.db.GetContext(ctx, &dbUser, u.selectUser, *filter.Email, *filter.Password)
	if err != nil && err.Error() == "sql: no rows in result set" {
		err = nil
	}
	user = storage.UserInfo(dbUser)
	return
}

func (u *User) SelectList(ctx context.Context, filter storage.UserFilter) (users []storage.UserInfo, err error) {
	if err = u.check(); err != nil {
		return
	}
	query := u.selectUserList

	if filter.ID != nil || filter.Email != nil {
		query += " WHERE"
	}

	if filter.ID != nil {
		query += fmt.Sprintf(" id = %d", *filter.ID)
	}

	if filter.ID != nil && filter.Email != nil {
		query += " OR"
	}

	if filter.Email != nil {
		query += fmt.Sprintf(` email = '%s'`, *filter.Email)
	}

	var dbUsers []userInfo
	err = u.db.SelectContext(ctx, &dbUsers, query)

	users = make([]storage.UserInfo, 0, len(dbUsers))
	for _, user := range dbUsers {
		users = append(users, storage.UserInfo(user))
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
