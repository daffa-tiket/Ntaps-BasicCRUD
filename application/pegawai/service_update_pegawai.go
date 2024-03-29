package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) UpdatePegawai(ctx context.Context, request pegawai_dto.UpdatePegawaiRequestDto) (*pegawai_dto.UpdatePegawaiResponseDto, error) {
	obj, err := s.RepositoryHolder.PegawaiRepository.FindOne(s.SharedHolder.Sql, "id = ?", request.ID)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `FindOne`")
		return nil, err
	}
	pegawai := createPegawaiFormRequest(request.Pegawai)
	obj.Nama = pegawai.Nama
	obj.Alamat = pegawai.Alamat
	obj.Telepon = pegawai.Telepon

	err = s.RepositoryHolder.PegawaiRepository.Update(s.SharedHolder.Sql, obj)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `Update`")
		return nil, err
	}

	return &pegawai_dto.UpdatePegawaiResponseDto{
		MandatoryRequestV2Dto: request.MandatoryRequestV2Dto,
		Pegawai:               obj,
	}, err

}
