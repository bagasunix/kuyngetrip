package domains

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories"
	"github.com/bagasunix/kuyngetrip/server/domains/usecases"
)

type Service interface {
	usecases.Authentication
}

type service struct {
	usecases.Authentication
}

// Builder Object for service
type serviceBuilder struct {
	logger       *zap.Logger
	jwtKey       string
	repositories repositories.Repositories
	middlewares  []Middleware
	redis        *redis.Client
}

// Constructor for serviceBuilder
func NewserviceBuilder(logs *zap.Logger, jwtKey string, repositories repositories.Repositories, redis *redis.Client) *serviceBuilder {
	o := new(serviceBuilder)
	o.logger = logs
	o.jwtKey = jwtKey
	o.repositories = repositories
	o.redis = redis
	return o
}

func buildService(logs *zap.Logger, jwtKey string, repo repositories.Repositories, redis *redis.Client) Service {
	svc := new(service)
	svc.Authentication = usecases.NewUser(logs, jwtKey, repo, redis)
	return svc
}

// Build Method which creates service
func (b *serviceBuilder) Build() Service {
	svc := buildService(b.logger, b.jwtKey, b.repositories, b.redis)
	for _, m := range b.middlewares {
		svc = m(svc)
	}
	return svc
}

// Setter method for the field middlewares of type []Middleware in the object serviceBuilder
func (s *serviceBuilder) SetMiddlewares(middlewares []Middleware) *serviceBuilder {
	s.middlewares = middlewares
	return s
}
