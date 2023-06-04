package payment

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Commond interface {
	Create(ctx context.Context, m models.Payment) error
	Updates(ctx context.Context, m models.Payment) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Payment])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Payment])
}

type Repository interface {
	Commond
	Query
	base.Repository
}
