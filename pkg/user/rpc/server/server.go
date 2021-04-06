package server

import (
	"context"

	"github.com/geoirb/todos/pkg/user"
	"github.com/geoirb/todos/pkg/user/rpc"
)

// AuthServer rpc for auth service
type AuthServer struct {
	// todo what is it?
	rpc.UnimplementedAuthServer
	svc *user.Service
}

// Authorization token and return user data
func (s *AuthServer) Authorization(ctx context.Context, req *rpc.Request) (res *rpc.Response, err error) {
	id, err := s.svc.Authorization(ctx, req.Token)
	res = &rpc.Response{
		Id: int32(id),
	}
	return
}

// NewAuthRPCServer ...
func NewAuthRPCServer(
	svc *user.Service,
) *AuthServer {
	return &AuthServer{
		svc: svc,
	}
}
