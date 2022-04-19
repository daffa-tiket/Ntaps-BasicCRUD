// Code generated by ntaps. DO NOT EDIT.
package di

import (
	"github.com/pkg/errors"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/logs"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/echokit"
)

func NewEchoLoggerWrapper(logger logs.Logger) (*echokit.LoggerWrapper, error) {
	return echokit.NewLoggerWrapper(logger)
}

func init() {
	if err := Container.Provide(NewEchoLoggerWrapper); err != nil {
		panic(errors.Wrap(err, "failed to provide echo-logger-wrapper"))
	}
}