package userdetails

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Command interface {
	Create(ctx context.Context, user *models.UserDetail) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, m *models.UserDetail) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.UserDetail])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserDetail])
	GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.UserDetail])
}

type Repository interface {
	Command
	Query
	base.Repository
}
