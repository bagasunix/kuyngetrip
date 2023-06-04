package userdetails

import (
	"context"
	"fmt"

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
func (g *gormProvider) Create(ctx context.Context, user *models.UserDetail) error {
	return err.ErrDuplicateValue(g.logs, g.GetModelName(), g.db.WithContext(ctx).Create(user).Error)
}

// Delete implements Repository.
func (g *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return err.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Delete(models.NewUserDetailBuilder().Build(), "id = ?", id.String()).Error)
}

// GetAll implements Repository.
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.UserDetail]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.UserDetail]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetByName implements Repository.
func (g *gormProvider) GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.UserDetail]) {
	a := fmt.Sprint('%', name, '%')
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("full_name like ?", a).Limit(int(limit)).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "User Details"
}

// Update implements Repository.
func (g *gormProvider) Update(ctx context.Context, m *models.UserDetail) error {
	return err.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Updates(m).Error)
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
