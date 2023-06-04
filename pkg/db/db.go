package db

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/bagasunix/kuyngetrip/pkg/err"
)

func NewDB(ctx context.Context, logger *zap.Logger, dbConfig *DbConfig) *gorm.DB {
	db, errs := gorm.Open(postgres.Open(dbConfig.GetDSN()), &gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	err.HandlerWithOSExit(logger, errs, "init", "database", "config", dbConfig.GetDSN())
	err.HandlerWithOSExit(logger, db.WithContext(ctx).Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(dbConfig.MaxIdleConns).SetMaxOpenConns(dbConfig.MaxOpenConns).SetConnMaxLifetime(time.Hour)), "db_resolver")
	return db
}
