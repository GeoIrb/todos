package client

import (
	"context"
)

// Auth rpc client.
type Auth interface {
	Authorization(ctx context.Context, token string) (id int, err error)
}
