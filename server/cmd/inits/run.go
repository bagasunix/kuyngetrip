package inits

import (
	"context"
	"flag"

	"github.com/bagasunix/kuyngetrip/pkg/env"
	"github.com/bagasunix/kuyngetrip/server/domains/data/repositories"
)

func Run() {
	flag.Parse()
	ctx := context.Background()
	configs, err := env.LoadEnv()
	logger := InitLogger()
	if err != nil {
		logger.Fatal(err.Error())
	}
	redis := InitRedis(ctx, configs)
	// ************ Database ************
	db := InitDb(ctx, logger, configs)
	Migrate(logger, db)
	// ************ Service ************
	repo := repositories.New(logger, db)
}
