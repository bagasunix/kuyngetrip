package userdetails

import (
	"context"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
)

type gormProvider struct {
	db   *gorm.DB
	logs *zap.Logger
}

// Create implements Repository.
func (g *gormProvider) Create(ctx context.Context, user *models.UserDetail) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetAll implements Repository.
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.UserDetail]) {
	panic("unimplemented")
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserDetail]) {
	panic("unimplemented")
}

// GetByName implements Repository.
func (g *gormProvider) GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.UserDetail]) {
	panic("unimplemented")
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	panic("unimplemented")
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	panic("unimplemented")
}

// Update implements Repository.
func (g *gormProvider) Update(ctx context.Context, m *models.UserDetail) error {
	panic("unimplemented")
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
