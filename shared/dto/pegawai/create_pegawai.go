package pegawai

import (
	dto "github.com/tiket/learn-crud/shared/dto"
)

type (
	CreatePegawaiRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		Pegawai               Pegawai                   `json:"pegawai"`
	}

	CreatePegawaiResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		Pegawai               interface{}
	}
)
