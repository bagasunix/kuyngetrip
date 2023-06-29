package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bagasunix/kuyngetrip/pkg/err"
	"github.com/bagasunix/kuyngetrip/server/endpoints/responses"
)

// EncodeError encode err from business-logic
func EncodeError(_ context.Context, errs error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	responseBuilder := responses.NewErrorBuilder()

	if strings.Contains(errs.Error(), err.ERR_NOT_AUTHORIZED) {
		w.WriteHeader(http.StatusUnauthorized)
		responseBuilder.SetError(errs)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}
	if strings.Contains(errs.Error(), err.ERR_INVALID_KEY) {
		w.WriteHeader(http.StatusBadRequest)
		responseBuilder.SetError(errs)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(errs.Error(), err.ERR_NOT_FOUND) {
		w.WriteHeader(http.StatusNotFound)
		responseBuilder.SetError(errs)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	if strings.Contains(errs.Error(), err.ERR_ALREADY_EXISTS) {
		w.WriteHeader(http.StatusConflict)
		responseBuilder.SetError(errs)
		json.NewEncoder(w).Encode(responseBuilder.Build())
		return
	}

	err.HandlerReturnedVoid(errs)
	w.WriteHeader(http.StatusInternalServerError)
	responseBuilder.SetError(err.CustomError(err.ERR_SOMETHING_WRONG))
	json.NewEncoder(w).Encode(responseBuilder.Build())
}
