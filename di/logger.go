// Code generated by ntaps. DO NOT EDIT.
package di

import (
	"github.com/pkg/errors"

	logs "github.com/tiket/TIX-HOTEL-UTILITIES-GO/logs"

	config "github.com/tiket/learn-crud/shared/config"
)

func NewLogger(conf *config.EnvConfiguration) (logs.Logger, error) {
	return logs.New(&logs.Option{
		LogFilePath: conf.LoggerFileName,
		Formatter:   logs.Formatter(conf.LoggerFormater),
		Level:       logs.Level(conf.LoggerLevel),
		MaxSize:     conf.LoggerMaxSize,
		MaxBackups:  conf.LoggerMaxBackups,
		MaxAge:      conf.LoggerMaxAge,
		Compress:    conf.LoggerCompress,
	})
}

func init() {
	if err := Container.Provide(NewLogger); err != nil {
		panic(errors.Wrap(err, "failed to provide logger"))
	}
}
