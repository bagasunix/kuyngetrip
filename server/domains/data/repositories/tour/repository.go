package tour

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Commond interface {
	Create(ctx context.Context, m *models.Tour) error
	Update(ctx context.Context, m *models.Tour) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Tour])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[models.Tour])
	GetByTour(ctx context.Context, tour string, limit int64) (result models.SliceResult[models.Tour])
}

type Repository interface {
	Commond
	Query
	base.Repository
}
