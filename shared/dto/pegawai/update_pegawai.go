package pegawai

import (
	dto "github.com/tiket/learn-crud/shared/dto"
)

type (
	UpdatePegawaiRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		Pegawai               Pegawai                   `json:"pegawai"`
		ID                    int
	}

	UpdatePegawaiResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		Pegawai               interface{}
	}
)
