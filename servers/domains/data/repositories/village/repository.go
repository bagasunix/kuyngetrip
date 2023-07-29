package village

import (
	"context"

	"bagasunix/kuyngetrip/servers/domains/data/models"
	"bagasunix/kuyngetrip/servers/domains/data/repositories/base"
)

type Query interface {
	GetById(ctx context.Context, id int64) (result models.SingleResult[*models.Village])
	GetByIdWithChannel(ctx context.Context, id int64, result chan models.SingleResult[*models.Village])
	GetByKeywords(ctx context.Context, limit int64, keywords string) (result models.SliceResult[models.Village], err error)
	GetByKeywordsByCoordinateByRadius(ctx context.Context, limit int64, keywords string, latitude float64, longitude float64, radius float32) (result models.SliceResult[models.Village], err error)
}

type Repository interface {
	base.Repository
	Query
}
