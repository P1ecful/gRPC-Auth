package authorization

import (
	"authorization/internal/db"
	"authorization/pkg/api"
	"context"
	"log"
	"strings"

	password "github.com/vzglad-smerti/password_hash"
)

type GRPCServer struct {
	api.UnimplementedAuthorizationServer
}

func (s *GRPCServer) Register(ctx context.Context, req *api.UserMeta) (*api.SuccessfulRegister, error) {
	hash, err := password.Hash(req.GetPassword())
	if err != nil {
		log.Fatal(err)
	}

	//FIXME!
	go db.AddMeta(req.GetLogin(), req.GetPassword())
	return &api.SuccessfulRegister{Login: strings.ToLower(req.GetLogin()),
		Password: hash}, nil
}
