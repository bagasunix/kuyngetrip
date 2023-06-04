package tourimage

import (
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/kuyngetrip/pkg/err"
	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
)

type gormProvider struct {
	db   *gorm.DB
	logs *zap.Logger
}

// Create implements Repository.
func (g *gormProvider) Create(ctx context.Context, m models.TourImage) error {
	return err.ErrDuplicateValue(g.logs, g.GetModelName(), g.db.WithContext(ctx).Create(&m).Error)
}

// Delete implements Repository.
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return err.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Delete(models.NewTourImageBuilder().Build(), "id = ?", id).Error)
}

// GetAll implements Repository.
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.TourImage]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.TourImage]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).First(&result.Value, "id = ?", id).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "Tour image"
}

// Updates implements Repository.
func (g *gormProvider) Updates(ctx context.Context, m models.TourImage) error {
	return err.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Updates(&m).Error)
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
