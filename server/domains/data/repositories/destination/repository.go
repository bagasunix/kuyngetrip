package destination

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Commond interface {
	Create(ctx context.Context, m *models.Destination) error
	Update(ctx context.Context, m *models.Destination) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Destination])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Destination])
	GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.Destination])
}

type Repository interface {
	Commond
	Query
	base.Repository
}
