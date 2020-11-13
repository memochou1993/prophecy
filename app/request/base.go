package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Validate(e *echo.Echo) *echo.Echo {
	e.Validator = &Validator{validator: validator.New()}

	return e
}
