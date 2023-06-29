package company

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Commond interface {
	CreateTx(ctx context.Context, tx any, m *models.Company) error
	Update(ctx context.Context, m *models.Company) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Company])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Company])
	GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.Company])
}

type Repository interface {
	Commond
	Query
	base.Repository
}
