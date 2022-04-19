package pegawai

import (
	dto "github.com/tiket/learn-crud/shared/dto"
)

type (
	DeletePegawaiRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		ID                    int
	}

	DeletePegawaiResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}
)
