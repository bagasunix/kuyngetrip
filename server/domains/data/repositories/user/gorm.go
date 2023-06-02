package user

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

// Update implements Repository.
func (u *gormProvider) Update(ctx context.Context, m *models.User) error {
	return err.ErrSomethingWrong(u.logs, u.db.WithContext(ctx).Updates(m).Error)
}

// UpdateStatus implements Repository.
func (u *gormProvider) UpdateStatus(ctx context.Context, m *models.User) error {
	return err.ErrSomethingWrong(u.logs, u.db.WithContext(ctx).Updates(m).Error)
}

// Create implements Repository.
func (u *gormProvider) Create(ctx context.Context, user *models.User) error {
	return err.ErrDuplicateValue(u.logs, u.GetModelName(), u.db.WithContext(ctx).Create(user).Error)
}

// Delete implements Repository.
func (u *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return err.ErrSomethingWrong(u.logs, u.db.WithContext(ctx).Delete(models.NewUserBuilder().Build(), "id = ?", id.String()).Error)
}

// GetAll implements Repository.
func (u *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.User]) {
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetByEmail implements Repository.
func (u *gormProvider) GetByEmail(ctx context.Context, email string) (result models.SingleResult[*models.User]) {
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).Where("email = ?", email).First(&result.Value).Error)
	return result
}

// GetById implements Repository.
func (u *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.User]) {
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).First(&result.Value, id).Error)
	return result
}

// GetByKeywords implements Repository.
func (u *gormProvider) GetByKeywords(ctx context.Context, keywords string, limit int64) (result models.SliceResult[models.User]) {
	a := fmt.Sprint('%', keywords, '%')
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).Where("email like ?", a).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetConnection implements Repository.
func (u *gormProvider) GetConnection() (T any) {
	return u.db
}

// GetModelName implements Repository.
func (u *gormProvider) GetModelName() string {
	return "User"
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
