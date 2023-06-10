package middlewares

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/server/endpoints"
)

func Logging(logger *zap.Logger) endpoints.Middleware {
	return func(e endpoints.Endpoint) endpoints.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Info("transport_error",
					zap.Error(err),
					zap.Duration("took", time.Since(begin)),
				)
			}(time.Now())
			return e(ctx, request)
		}
	}
}
