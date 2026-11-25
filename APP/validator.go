package app

import "github.com/go-playground/validator/v10"

func RedValidator() *validator.Validate {
	return validator.New()
}
