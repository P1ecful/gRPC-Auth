package authorization

import (
	"authorization/internal/db"
	"authorization/internal/entity"
	"authorization/pkg/api"
	"context"
)

type GRPCServer struct {
	api.UnimplementedAuthorizationServer
}

func (s *GRPCServer) Register(ctx context.Context, req *api.UserMeta) (*api.SuccessfulResponse, error) {
	hashedbymd5 := entity.GetMd5(req.GetLogin(), req.GetPassword())
	newUser := db.User{UserHash: hashedbymd5}

	return &api.SuccessfulResponse{Status: newUser.Register()}, nil // server response
}

func (s *GRPCServer) Auth(ctx context.Context, req *api.UserMeta) (*api.SuccessfulResponse, error) {

	hashedbymd5 := entity.GetMd5(req.GetLogin(), req.GetPassword())
	newUser := db.User{UserHash: hashedbymd5}

	return &api.SuccessfulResponse{Status: newUser.Auth()}, nil // server response
}
