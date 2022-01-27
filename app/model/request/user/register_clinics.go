package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegRegister struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	ServiceTime string `json:"service_time" form:"service_time"`
	City        string `json:"city" form:"city"`
	Contact     string `json:"contact" form:"contact"`
	Province    string `json:"province" form:"province"`
	PostalCode  string `json:"postal_code" form:"postal_code"`
}

func (reg RegRegister) Validate() error {
	return validation.ValidateStruct(&reg,
		validation.Field(&reg.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&reg.Address, validation.Required, validation.Length(6, 100)),
		validation.Field(&reg.ServiceTime, validation.Required, validation.Length(6, 100)),
		validation.Field(&reg.City, validation.Required, validation.Length(6, 100)),
		validation.Field(&reg.Contact, validation.Required, validation.Length(11, 14)),
		validation.Field(&reg.Province, validation.Required, validation.Length(3, 20)),
		validation.Field(&reg.PostalCode, validation.Required, validation.Length(4, 10)),
	)
}
