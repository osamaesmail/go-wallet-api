package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"reflect"
)

type Validation struct {
}

func NewValidation() (validate *validator.Validate, err error) {
	validate = validator.New()
	err = decimalGreaterThan(validate)
	return
}

func decimalGreaterThan(validate *validator.Validate) error {
	validate.RegisterCustomTypeFunc(
		func(field reflect.Value) interface{} {
			if valuer, ok := field.Interface().(decimal.Decimal); ok {
				return valuer.String()
			}
			return nil
		}, decimal.Decimal{},
	)
	if err := validate.RegisterValidation(
		"dge", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.GreaterThanOrEqual(baseValue)
		},
	); err != nil {
		return err
	}
	
	return nil
}
