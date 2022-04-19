// Code generated by ntaps. DO NOT EDIT.
package shared

import (
	"go.uber.org/dig"
	"log"

	"io"

	persistent "github.com/tiket/TIX-HOTEL-UTILITIES-GO/persistent"

	logs "github.com/tiket/TIX-HOTEL-UTILITIES-GO/logs"

	migrator "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/migration"

	echo "github.com/labstack/echo/v4"

	echokit "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/echokit"

	validator "github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/validator"

	"github.com/tiket/learn-crud/shared/config"
)

type (
	Holder struct {
		dig.In
		Config *config.EnvConfiguration

		Sql persistent.ORM

		Logger logs.Logger

		Migration migrator.Tool

		Echo *echo.Echo

		EchoLoggerWrapper *echokit.LoggerWrapper

		Validator validator.Validator

		CustomHolder CustomHolder
	}
)

func (h *Holder) Close() {
	log.Print("closing resource")

	var i interface{} = nil

	i = h.Sql
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Logger
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Migration
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Echo
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.EchoLoggerWrapper
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Validator
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	h.CustomHolder.Close()

	log.Print("done closing resource")
}