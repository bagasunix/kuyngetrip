package middlewares

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/server/domains"
	"github.com/bagasunix/kuyngetrip/server/endpoints/requests"
	"github.com/bagasunix/kuyngetrip/server/endpoints/responses"
)

type loggingMiddleware struct {
	logs *zap.Logger
	next domains.Service
}

// CreateUser implements domains.Service.
func (l *loggingMiddleware) CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, errs error) {
	defer func(begin time.Time) {
		l.logs.Info("CreateUser", zap.String("request", string(request.ToJSON())), zap.String("response", string(response.ToJSON())), zap.Error(errs), zap.Duration("took", time.Since(begin)))
	}(time.Now())
	return l.next.CreateUser(ctx, request)
}

func LoggingMiddleware(logs *zap.Logger) domains.Middleware {
	return func(next domains.Service) domains.Service {
		return &loggingMiddleware{logs: logs, next: next}
	}
}
