package inits

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bagasunix/kuyngetrip/pkg/err"
	"github.com/bagasunix/kuyngetrip/server/domains/data/models"
)

func GetTables() (tables []interface{}) {
	tables = append(tables, models.NewCountryBuilder().Build())
	tables = append(tables, models.NewProvinceBuilder().Build())
	tables = append(tables, models.NewCityBuilder().Build())
	tables = append(tables, models.NewSubDistrictBuilder().Build())
	tables = append(tables, models.NewVillageBuilder().Build())
	tables = append(tables, models.NewCoordinateBuilder().Build())

	tables = append(tables, models.NewUserBuilder().Build())
	tables = append(tables, models.NewUserDetailBuilder().Build())

	tables = append(tables, models.NewTourBuilder().Build())
	tables = append(tables, models.NewDestinationBuilder().Build())
	tables = append(tables, models.NewTourImageBuilder().Build())
	tables = append(tables, models.NewTourScheduleBuilder().Build())
	tables = append(tables, models.NewParticipantBuilder().Build())
	tables = append(tables, models.NewTourReviewBuilder().Build())
	tables = append(tables, models.NewPaymentBuilder().Build())

	return tables
}

func Migrate(logs *zap.Logger, db *gorm.DB) {
	err.HandlerWithOSExit(logs, db.AutoMigrate(GetTables()...), "AutoMigrate")
}
