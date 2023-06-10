package endpoints

import (
	"context"

	"github.com/bagasunix/kuyngetrip/server/domains"
)

// Endpoint is the fundamental building block of servers and clients.
// It represents a single RPC method.
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Middleware is a chainable behavior modifier for endpoints.
type Middleware func(Endpoint) Endpoint

type Endpoints struct {
	UserEndpoint UserEndpoint
}

// Builder EndpointsBuilder Builder Object for Endpoints
type Builder struct {
	service     domains.Service
	middlewares map[string][]Middleware
}

// NewBuilder Constructor for EndpointsBuilder
func NewBuilder() *Builder {
	o := new(Builder)
	return o
}

// Build Method which creates Endpoints
func (b *Builder) Build() Endpoints {
	eps := new(Endpoints)
	eps.UserEndpoint = NewUserEndpoint(b.service, b.middlewares)

	return *eps
}

// SetService Setter method for the field service of type domains.Service in the object EndpointsBuilder
func (b *Builder) SetService(service domains.Service) *Builder {
	b.service = service
	return b
}

// SetMiddlewares Setter method for the field middlewares of type map[string][]endpoint.Middleware in the object EndpointsBuilder
func (b *Builder) SetMiddlewares(middlewares map[string][]Middleware) *Builder {
	b.middlewares = middlewares
	return b
}
