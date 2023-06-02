package user

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Command interface {
	Create(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, m *models.User) error
	UpdateStatus(ctx context.Context, m *models.User) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.User])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.User])
	GetByEmail(ctx context.Context, email string) (result models.SingleResult[*models.User])
	GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.User])
}

type Repository interface {
	Command
	Query
	base.Repository
}
