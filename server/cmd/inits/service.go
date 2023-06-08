package inits

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/pkg/env"
	"github.com/bagasunix/kuyngetrip/server/domains"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories"
	"github.com/bagasunix/kuyngetrip/server/domains/middlewares"
)

func InitService(logger *zap.Logger, configs *env.Configs, repositories repositories.Repositories, redis *redis.Client) domains.Service {
	serviceBuilder := domains.NewserviceBuilder(logger, configs.JwtSecret, repositories, redis)
	serviceBuilder.SetMiddlewares(getServiceMiddleware(logger))
	return serviceBuilder.Build()
}

func getServiceMiddleware(logger *zap.Logger) []domains.Middleware {
	var mw []domains.Middleware
	mw = addDefaultServiceMiddleware(logger, mw)
	return mw
}

func addDefaultServiceMiddleware(logger *zap.Logger, mw []domains.Middleware) []domains.Middleware {
	return append(mw, middlewares.LoggingMiddleware(logger))
}
