package pegawai

import (
	"go.uber.org/dig"
)

type RepositoryHolder struct {
	dig.In

	// - -model-start
	PegawaiRepository PegawaiRepository
	// - -model-end

}
