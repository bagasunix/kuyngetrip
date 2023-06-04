package tourschedule

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories/base"
)

type Commond interface {
	Create(ctx context.Context, m models.TourSchedule) error
	Updates(ctx context.Context, m models.TourSchedule) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type Query interface {
	GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.TourSchedule])
	GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.TourSchedule])
}

type Repository interface {
	Commond
	Query
	base.Repository
}
