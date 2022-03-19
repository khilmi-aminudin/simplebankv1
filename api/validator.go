package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/khilmi-aminudin/simplebankv1/utils"
)

// validator untuk currency dengan menekstend validator.Func
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
