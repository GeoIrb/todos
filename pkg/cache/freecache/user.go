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
func (c *User) SetPassword(ctx context.Context, mail, password string, ttl time.Duration) (err error) {
	return c.cache.Set(
		[]byte(mail),
		[]byte(password),
		int(ttl.Seconds()),
	)
}

// GetPassword by email.
func (c *User) GetPassword(ctx context.Context, mail string) (password string, isExist bool, err error) {
	var value []byte
	if value, err = c.cache.Get([]byte(mail)); err != nil || value == nil {
		isExist = false
		if err == freecache.ErrNotFound {
			err = nil
		}
		return
	}
	password = string(value)
	return
}
