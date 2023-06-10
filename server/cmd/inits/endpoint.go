package inits

import (
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/server/domains"
	"github.com/bagasunix/kuyngetrip/server/endpoints"
	"github.com/bagasunix/kuyngetrip/server/endpoints/middlewares"
)

func InitEndpoint(logger *zap.Logger, s domains.Service) endpoints.Endpoints {
	a := endpoints.NewBuilder()
	a.SetMiddlewares(getEndpointMiddleware(logger))
	a.SetService(s)
	return a.Build()
}
func getEndpointMiddleware(logger *zap.Logger) (mw map[string][]endpoints.Middleware) {
	mw = map[string][]endpoints.Middleware{}
	addDefaultEndpointMiddleware(logger, mw)
	return mw
}
func defaultMiddlewares(logger *zap.Logger, method string) []endpoints.Middleware {
	return []endpoints.Middleware{
		middlewares.Logging(logger.With(zap.String("method", method))),
	}
}
func middlewaresWithAuthentication(logger *zap.Logger, method string) []endpoints.Middleware {
	mw := defaultMiddlewares(logger, method)
	//return mw
	return append(mw, middlewares.Authentication())
}
func addDefaultEndpointMiddleware(logger *zap.Logger, mw map[string][]endpoints.Middleware) {
	mw[endpoints.CREATE_USER] = middlewaresWithAuthentication(logger, endpoints.CREATE_USER)

}
