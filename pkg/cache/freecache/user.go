package freecache

import (
	"context"
	"time"

	"github.com/coocood/freecache"

	"github.com/geoirb/todos/pkg/cache"
)

// User memory cache.
type User struct {
	cache *freecache.Cache
}

var _ cache.User = &User{}

// New returns memory freecache for user.
func NewUser(
	size int,
) *User {
	return &User{
		cache: freecache.NewCache(size),
	}
}

// SetPassword saves temporary password for new user.
func (u *User) SetPassword(ctx context.Context, email, password string, ttl time.Duration) (err error) {
	return u.cache.Set(
		[]byte(email),
		[]byte(password),
		int(ttl),
	)
}

// GetPassword by email.
func (u *User) GetPassword(ctx context.Context, email string) (password string, isExist bool, err error) {
	var value []byte
	if value, err = u.cache.Get([]byte(email)); err != nil {
		if err == freecache.ErrNotFound {
			err = nil
		}
		return
	}
	isExist = true
	password = string(value)
	return
}

func (u *User) DeletePassword(ctx context.Context, email string) {
	u.cache.Del([]byte(email))
}
