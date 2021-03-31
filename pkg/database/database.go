package database

import (
	"context"
)

// User database.
type User interface {
	Insert(ctx context.Context, user UserInfo) error
	Select(ctx context.Context, filter UserFilter) (UserInfo, error)
}
