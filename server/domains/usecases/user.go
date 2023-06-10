package usecases

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/pkg/err"
	"github.com/bagasunix/kuyngetrip/pkg/helpers"
	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories"
	"github.com/bagasunix/kuyngetrip/server/endpoints/requests"
	"github.com/bagasunix/kuyngetrip/server/endpoints/responses"
)

type UserWithEndpoint interface {
	CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, errs error)
}
type UserWithNoEndpoint interface {
}

type Authentication interface {
	UserWithEndpoint
	UserWithNoEndpoint
}

type authentication struct {
	repo   repositories.Repositories
	jwtKey string
	logger *zap.Logger
	redis  *redis.Client
}

// CreateUser implements Authentication.
func (a *authentication) CreateUser(ctx context.Context, request *requests.CreateUser) (response *responses.EntityId, errs error) {
	resBuilder := responses.NewEntityIdBuilder()

	if errs = request.Validate(); errs != nil {
		return resBuilder.Build(), errs
	}

	if helpers.IsEmailValid(request.Email) != true {
		return resBuilder.Build(), err.ErrValidEmail(a.logger, request.Email)
	}

	mUser := models.NewUserBuilder()
	mUser.SetId(helpers.GenerateUUIDV6(a.logger))
	mUser.SetEmail(request.Email)
	mUser.SetPassword(helpers.HashAndSalt([]byte(request.Password)))
	mUser.SetUserType(request.UserType)
	mUser.SetUserName(request.UserName)
	mUser.SetActivite("1")

	if errs = a.repo.GetUser().Create(ctx, mUser.Build()); errs != nil {
		return resBuilder.Build(), errs
	}

	return resBuilder.SetId(mUser.Build().ID).Build(), nil
}

func NewUser(logger *zap.Logger, jwtKey string, repo repositories.Repositories, redis *redis.Client) Authentication {
	a := new(authentication)
	a.logger = logger
	a.repo = repo
	a.jwtKey = jwtKey
	a.redis = redis
	return a
}
