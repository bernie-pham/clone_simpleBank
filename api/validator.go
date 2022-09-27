package api

import (
	"github.com/bernie-pham/cloneSimpleBank/ultilities"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return ultilities.IsSupportedCurrency(currency)
	}
	return false
}
