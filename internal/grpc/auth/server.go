package auth

import (
	"context"

	ssoy1 "github.com/aolychkin/protos/gen/go/sso"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	Login(ctx context.Context,
		email string,
		password string,
		appID int,
		) (token string, err error)
		
		RegisterNewUser (ctx context.Context,		
		email string,		
		password string,		
		) (userID int64, err error)
		
		IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	ssoy1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssoy1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}


// SERVISE LAYOUT handler
func (s *serverAPI) Login(
	ctx context.Context,
	req *ssoy1.LoginRequest,
) (*ssoy1.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))

	if err != nil{
		// TODO 
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssoy1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssoy1.RegisterRequest,
) (*ssoy1.RegisterResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	userID, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil{
		// TODO: ...
		return nil, status.Error(code.Insernal, "internal error")
	}	

	return  &ssoy1.RegisterResponse{
		UserID: userID,
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssoy1.IsAdminRequest,
) (*ssoy1.IsAdminResponse, error) {
	if err := validateIsAmin(req); err != nil {
		return nil, err
	}

	IsAdmin, err := s.auth.IsAdmin(ctx, req.GetUserIdl())
	if err != nil{
		// TODO: ...
		return nil, status.Error(code.Insernal, "internal error")
	}	

	return  &ssoy1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

// Validation functions
func validateLogin(req *ssoy1.LoginRequest) error{
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.GetAppId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	return nil
}

func validateRegister(req *ssoy1.RegisterRequest) error{
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	return nil
}

func validateIsAmin(req *ssoy1.IsAdminRequest) error{
	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	return nil
}