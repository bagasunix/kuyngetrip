package endpoints

import (
	"context"

	"github.com/bagasunix/kuyngetrip/server/domains"
	"github.com/bagasunix/kuyngetrip/server/endpoints/requests"
)

const (
	CREATE_USER = "CreateUser"
)

type UserEndpoint struct {
	CreateUserEndpoint Endpoint
}

func NewUserEndpoint(s domains.Service, mdw map[string][]Middleware) UserEndpoint {
	eps := UserEndpoint{}
	eps.CreateUserEndpoint = makeCreateUserEndpoint(s)
	SetEndpoint(CREATE_USER, &eps.CreateUserEndpoint, mdw)
	return eps
}

func makeCreateUserEndpoint(s domains.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return s.CreateUser(ctx, request.(*requests.CreateUser))
	}
}
