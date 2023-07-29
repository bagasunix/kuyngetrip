package errors

import "errors"

var (
	ErrInvalidActivityId       = errors.New("invalid activity id")
	ErrInvalidActivityType     = errors.New("invalid activity type")
	ErrInvalidActivityDate     = errors.New("invalid activity date")
	ErrInvalidCompanyId        = errors.New("invalid companyId")
	ErrInvalidActivityCategory = errors.New("invalid activity category")
)
