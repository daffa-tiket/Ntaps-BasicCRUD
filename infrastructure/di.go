// Code generated by ntaps. DO NOT EDIT.
package infrastructure

import (
	"github.com/pkg/errors"

	pegawai_http "github.com/tiket/learn-crud/infrastructure/http/pegawai"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {

	// - pegawai-http-start
	if err := container.Provide(pegawai_http.NewController); err != nil {
		return errors.Wrap(err, "failed to provide pegawai controller")
	}
	// - pegawai-http-end

	// - infrastructure-di-end
	return nil
}
