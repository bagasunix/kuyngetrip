package company

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

// CreateTx implements Repository.
func (g *gormProvider) CreateTx(ctx context.Context, tx any, m *models.Company) error {
	return err.ErrDuplicateValue(g.logs, g.GetModelName(), tx.(*gorm.DB).WithContext(ctx).Create(&m).Error)
}

// GetAll implements Repository.
func (g *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Company]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetById implements Repository.
func (g *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Company]) {
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("id = ?", id).First(&result.Value).Error)
	return result
}

// GetByName implements Repository.
func (g *gormProvider) GetByName(ctx context.Context, name string, limit int64) (result models.SliceResult[models.Company]) {
	companyName := fmt.Sprint('%', name, '%')
	result.Error = err.ErrRecordNotFound(g.logs, g.GetModelName(), g.db.WithContext(ctx).Where("company_name like ?", companyName).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (g *gormProvider) GetConnection() (T any) {
	return g.db
}

// GetModelName implements Repository.
func (g *gormProvider) GetModelName() string {
	return "Company"
}

// Update implements Repository.
func (g *gormProvider) Update(ctx context.Context, m *models.Company) error {
	return err.ErrSomethingWrong(g.logs, g.db.WithContext(ctx).Updates(&m).Error)
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
