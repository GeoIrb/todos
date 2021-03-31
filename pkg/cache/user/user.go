package user

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

func (c *User) SetPassword(ctx context.Context, key, password string, ttl time.Duration) (err error) {
	return c.cache.Set(
		[]byte(key),
		[]byte(password),
		int(ttl.Seconds()),
	)
}

func (c *User) GetPassword(ctx context.Context, key string) (password string, isExist bool, err error) {
	var value []byte
	if value, err = c.cache.Get([]byte(key)); err != nil || value == nil {
		isExist = false
		if err == freecache.ErrNotFound {
			err = nil
		}
		return
	}
	password = string(value)
	return
}

// Delete user by key.
func (c *User) Delete(ctx context.Context, key string) (affected bool) {
	return c.cache.Del([]byte(key))
}

// New returns memory freecache for user.
func New(
	size int,
) *User {
	return &User{
		cache: freecache.NewCache(size),
	}
}
