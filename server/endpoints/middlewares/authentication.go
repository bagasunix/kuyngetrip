package middlewares

import (
	"context"
	"fmt"

	"github.com/bagasunix/kuyngetrip/pkg/err"
	"github.com/bagasunix/kuyngetrip/server/endpoints"
)

func Authentication() endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, errs error) {
			var s string
			var ok bool

			if s, ok = ctx.Value("JWTToken").(string); !ok {
				return nil, err.ErrUnAuthorized()
			}

			fmt.Print(s)
			return e(ctx, request)
		}
	}
}
