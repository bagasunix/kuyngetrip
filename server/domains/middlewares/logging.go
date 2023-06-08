package middlewares

import (
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/server/domains"
)

type loggingMiddleware struct {
	logs *zap.Logger
	next domains.Service
}

func LoggingMiddleware(logs *zap.Logger) domains.Middleware {
	return func(next domains.Service) domains.Service {
		return &loggingMiddleware{logs: logs, next: next}
	}
}
