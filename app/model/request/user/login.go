//Package request provide object struct and validation for usecase if want to store data to database
package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegLoginEmail struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (rg RegLoginEmail) Validate() error {
	return validation.ValidateStruct(&rg,
		validation.Field(&rg.Email, validation.Required, validation.Length(1, 100)),
		validation.Field(&rg.Password, validation.Required, validation.Length(1, 100)),
	)
}
