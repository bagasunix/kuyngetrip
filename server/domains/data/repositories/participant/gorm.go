package participant

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

// Update implements Repository.
func (u *gormProvider) Update(ctx context.Context, m *models.Participant) error {
	return err.ErrSomethingWrong(u.logs, u.db.WithContext(ctx).Updates(m).Error)
}

// Create implements Repository.
func (u *gormProvider) Create(ctx context.Context, user *models.Participant) error {
	return err.ErrDuplicateValue(u.logs, u.GetModelName(), u.db.WithContext(ctx).Create(user).Error)
}

// Delete implements Repository.
func (u *gormProvider) Delete(ctx context.Context, id uuid.UUID) error {
	return err.ErrSomethingWrong(u.logs, u.db.WithContext(ctx).Delete(models.NewUserBuilder().Build(), "id = ?", id.String()).Error)
}

// GetAll implements Repository.
func (u *gormProvider) GetAll(ctx context.Context, limit int64) (result models.SliceResult[models.Participant]) {
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).Limit(int(limit)).Find(&result.Value).Error)
	return result
}

// GetById implements Repository.
func (u *gormProvider) GetById(ctx context.Context, id uuid.UUID) (result models.SingleResult[*models.Participant]) {
	result.Error = err.ErrRecordNotFound(u.logs, u.GetModelName(), u.db.WithContext(ctx).First(&result.Value, id).Error)
	return result
}

// GetConnection implements Repository.
func (u *gormProvider) GetConnection() (T any) {
	return u.db
}

// GetModelName implements Repository.
func (u *gormProvider) GetModelName() string {
	return "Participant"
}

func NewGorm(logs *zap.Logger, db *gorm.DB) Repository {
	g := new(gormProvider)
	g.logs = logs
	g.db = db
	return g
}
