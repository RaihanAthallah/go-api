package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegScreening struct {
	Number1 string `json:"number1" form:"number1"`
	Number2 string `json:"number2" form:"number2"`
	Number3 string `json:"number3" form:"number3"`
	Number4 string `json:"number4" form:"number4"`
	Number5 string `json:"number5" form:"number5"`
}

func (reg RegScreening) Validate() error {
	return validation.ValidateStruct(&reg,
		validation.Field(&reg.Number1, validation.Required, validation.Length(1, 5)),
		validation.Field(&reg.Number2, validation.Required, validation.Length(1, 5)),
		validation.Field(&reg.Number3, validation.Required, validation.Length(1, 5)),
		validation.Field(&reg.Number4, validation.Required, validation.Length(1, 5)),
		validation.Field(&reg.Number5, validation.Required, validation.Length(1, 5)),
	)
}
