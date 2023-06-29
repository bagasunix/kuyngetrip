package endpoints

import "context"

// Endpoint is the fundamental building block of servers and clients.
// It represents a single RPC method.
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Middleware is a chainable behavior modifier for endpoints.
type Middleware func(Endpoint) Endpoint

func SetEndpoint(endpointName string, endpoint *Endpoint, mdw map[string][]Middleware) {
	for _, m := range mdw[endpointName] {
		*endpoint = m(*endpoint)
	}
}
