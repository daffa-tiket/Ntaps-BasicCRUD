package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) GetAllPegawai(ctx context.Context, request pegawai_dto.GetAllPegawaiRequestDto) (*pegawai_dto.GetAllPegawaiResponseDto, error) {
	value, err := s.RepositoryHolder.PegawaiRepository.Find(s.SharedHolder.Sql)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `Find`")
		return nil, err
	}
	return &pegawai_dto.GetAllPegawaiResponseDto{
		MandatoryRequestV2Dto: request.MandatoryRequestV2Dto,
		Data:                  value,
	}, nil

}
