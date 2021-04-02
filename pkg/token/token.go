package token

import (
	"context"
)

// Token for storage in context.
type Token struct {
	field interface{}
}

// New ...
func New() *Token {
	return &Token{
		field: "token",
	}
}

// Put token in context.
func (t *Token) Put(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, t.field, token)
}

// Get token from context.
func (t *Token) Get(ctx context.Context) string {
	token, _ := ctx.Value(t.field).(string)
	return token
}
