// Code generated by ntaps. DO NOT EDIT.
package interfaces

import (
	"go.uber.org/dig"

	"github.com/tiket/learn-crud/interfaces/pegawai"
)

type (
	Holder struct {
		dig.In

		// - Pegawai-view-service-start
		PegawaiViewService pegawai.ViewService
		// - Pegawai-view-service-end

		// interfaces-view-service-end
	}
)
