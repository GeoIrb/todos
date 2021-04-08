package user

import (
	"context"
)

type Service interface {
	New(ctx context.Context, info Registartion) error
	Login(ctx context.Context, info Login) (auth Auth, err error)
	Create(ctx context.Context, info Create) (err error)
	Authorization(ctx context.Context, token string) (id int, err error)
	GetUserList(ctx context.Context, filter Filter) (users []UserInfo, err error)
}
