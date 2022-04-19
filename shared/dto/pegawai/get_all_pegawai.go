package pegawai

import (
	dto "github.com/tiket/learn-crud/shared/dto"
)

type (
	GetAllPegawaiRequestDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
	}

	GetAllPegawaiResponseDto struct {
		MandatoryRequestV2Dto dto.MandatoryRequestV2Dto `json:"-"`
		Data                  interface{}
	}
)
