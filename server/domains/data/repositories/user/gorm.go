package user

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
func (u *gormProvider) Create(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}

// Delete implements Repository.
func (u *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetAll implements Repository.
func (u *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.User]) {
	panic("unimplemented")
}

// GetByEmail implements Repository.
func (u *gormProvider) GetByEmail(ctx context.Context, email string) (result models.SingleResult[*models.User]) {
	panic("unimplemented")
}

// GetById implements Repository.
func (u *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.User]) {
	panic("unimplemented")
}

// GetByKeywords implements Repository.
func (u *gormProvider) GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.User]) {
	panic("unimplemented")
}

// GetConnection implements Repository.
func (u *gormProvider) GetConnection() (T any) {
	panic("unimplemented")
}

// GetModelName implements Repository.
func (u *gormProvider) GetModelName() string {
	panic("unimplemented")
}

// UpdateStatus implements Repository.
func (u *gormProvider) UpdateStatus(ctx context.Context, model *models.User) error {
	panic("unimplemented")
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
