package requests

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/bagasunix/kuyngetrip/pkg/err"
)

type CreateUser struct {
	UserName             string  `json:"user_name"`
	Email                string  `json:"email"`
	Password             string  `json:"password"`
	PasswordConfirmation string  `json:"password_confirmation"`
	UserType             string  `json:"user_type"`
	Picture              *string `json:"picture"`
	Active               string  `json:"active"`
}

func (s *CreateUser) Validate() error {
	if validation.IsEmpty(s.UserName) {
		return err.ErrInvalidAttributes("username")
	}
	if validation.IsEmpty(s.Email) {
		return err.ErrInvalidAttributes("email")
	}
	if validation.IsEmpty(s.Password) {
		return err.ErrInvalidAttributes("password")
	}
	if validation.IsEmpty(s.PasswordConfirmation) {
		return err.ErrInvalidAttributes("password confirmation")
	}
	if s.Password != s.PasswordConfirmation {
		return err.CustomError("password not the same as password")
	}
	if validation.IsEmpty(s.UserType) {
		return err.ErrInvalidAttributes("usertype")
	}
	if validation.IsEmpty(s.Active) {
		return err.ErrInvalidAttributes("status")
	}
	return nil
}

func (s *CreateUser) ToJSON() []byte {
	j, errs := json.Marshal(s)
	err.HandlerReturnedVoid(errs)
	return j
}
