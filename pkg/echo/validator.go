package echo

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func NewValidate() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
