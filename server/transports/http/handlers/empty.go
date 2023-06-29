package handlers

import (
	"context"
	"net/http"

	"github.com/bagasunix/kuyngetrip/server/endpoints/requests"
)

func decodeEmpty(ctx context.Context, request2 *http.Request) (request interface{}, errs error) {
	return new(requests.Empty), nil
}
