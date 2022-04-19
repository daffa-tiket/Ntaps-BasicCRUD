package pegawai

import (
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"time"
)

// - private function & utility for pegawai goes here

func createPegawaiFormRequest(request pegawai_dto.Pegawai) Pegawai {
	pegawai := Pegawai{
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		Telepon: request.Telepon,
		Version: request.Version,
	}

	pegawai.IsDeleted = false

	if pegawai.CreatedDate.IsZero() {
		pegawai.CreatedDate = time.Now()
	}

	if pegawai.UpdatedDate.IsZero() {
		pegawai.UpdatedDate = time.Now()
	}

	//if pegawai.CreatedBy == "" {
	//	pegawai.CreatedBy = constant.System
	//}
	//
	//if pegawai.UpdatedBy == "" {
	//	pegawai.UpdatedBy = constant.System
	//}

	return pegawai
}
