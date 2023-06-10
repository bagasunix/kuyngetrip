package helpers

import (
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	"github.com/bagasunix/kuyngetrip/pkg/err"
)

func GenerateUUIDV1(logger *zap.Logger) uuid.UUID {
	id, errs := uuid.NewV1()
	err.HandlerWithLoggerReturnedVoid(logger, errs, "uuid", "generator")
	return id
}

func GenerateUUIDV4(logger *zap.Logger) uuid.UUID {
	id, errs := uuid.NewV4()
	err.HandlerWithLoggerReturnedVoid(logger, errs, "uuid", "generator")
	return id
}

func GenerateUUIDV6(logs *zap.Logger) uuid.UUID {
	id, errs := uuid.NewV6()
	err.HandlerWithLoggerReturnedVoid(logs, errs, "uuid", "generator")
	return id
}
