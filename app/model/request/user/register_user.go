package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegRegisterUser struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Handphone string `json:"handphone" form:"handphone"`
}

func (reg RegRegisterUser) Validate() error {
	return validation.ValidateStruct(&reg,
		validation.Field(&reg.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&reg.Email, validation.Required, is.Email),
		validation.Field(&reg.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&reg.Handphone, validation.Required, validation.Length(11, 14)),
	)
}
