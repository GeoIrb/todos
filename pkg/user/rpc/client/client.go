package client

import (
	"context"

	"github.com/geoirb/todos/pkg/user/rpc"
)

// AuthClient rpc for auth service.
type AuthClient struct {
	client rpc.AuthClient
}

var _ Auth = &AuthClient{}

// Authorization token and return user data.
func (c *AuthClient) Authorization(ctx context.Context, token string) (id string, err error) {
	if res, err := c.client.Authorization(
		ctx,
		&rpc.Request{
			Token: token,
		},
	); err == nil {
		id = res.Id
	}
	return
}

// NewAuthRPCClient ...
func NewAuthRPCClient(
	client rpc.AuthClient,
) *AuthClient {
	return &AuthClient{
		client: client,
	}
}
