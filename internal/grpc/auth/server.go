package auth

import (
	"context"

	"google.golang.org/grpc"
)

type serverAPI struct {
	ssoy1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssoy1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssoy1.LoginRequest,
) (*ssoy1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssoy1.LoginRequest,
) (*ssoy1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssoy1.LoginRequest,
) (*ssoy1.LoginResponse, error) {
	panic("implement me")
}
