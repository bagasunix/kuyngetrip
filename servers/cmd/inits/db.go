package inits

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"bagasunix/kuyngetrip/pkg/db"
	"bagasunix/kuyngetrip/pkg/env"
)

func InitDb(ctx context.Context, logger zap.Logger, configs *env.Configs) *gorm.DB {
	configBuilder := db.NewDbConfigBuilder()
	configBuilder.SetDriver(configs.DbDriver)
	configBuilder.SetHost(configs.DbHost)
	configBuilder.SetPort(configs.DbPort)
	configBuilder.SetUser(configs.DbUsername)
	configBuilder.SetDatabaseName(configs.DbName)
	configBuilder.SetPassword(configs.DbPassword)

	config := configBuilder.Build()
	return db.NewDB(ctx, logger, config)
}
