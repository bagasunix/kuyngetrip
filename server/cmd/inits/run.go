package inits

import (
	"context"
	"flag"

	"github.com/bagasunix/kuyngetrip/pkg/env"
)

func Run() {
	flag.Parse()
	ctx := context.Background()
	configs, err := env.LoadEnv()
	logger := InitLogger()
	if err != nil {
		logger.Fatal(err.Error())
	}
	// ************ Database ************
	db := InitDb(ctx, logger, configs)
	Migrate(logger, db)
}
