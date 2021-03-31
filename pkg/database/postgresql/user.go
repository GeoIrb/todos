package postgresql

import (
	"context"

	"github.com/geoirb/todos/pkg/database"
)

// User database.
type User struct{}

func (u *User) Insert(ctx context.Context, user database.UserInfo) error {
	return nil
}
func (u *User) Select(ctx context.Context, filter database.UserFilter) (database.UserInfo, error) {
	return database.UserInfo{}, nil
}

func NewUser() {}
