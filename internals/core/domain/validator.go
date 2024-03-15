package domain

import "github.com/go-playground/validator"

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator.New()}
}
