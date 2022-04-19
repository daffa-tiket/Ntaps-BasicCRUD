// Code generated by ntaps. DO NOT EDIT.
package di

import (
	"github.com/pkg/errors"

	echo "github.com/labstack/echo/v4"

	echokit "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/echokit"

	validator "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/validator"
)

func NewEcho(logger *echokit.LoggerWrapper, validator validator.Validator) (*echo.Echo, error) {
	e := echo.New()
	e.Logger = logger
	e.Validator = validator
	return e, nil
}

func init() {
	if err := Container.Provide(NewEcho); err != nil {
		panic(errors.Wrap(err, "failed to provide echo"))
	}
}