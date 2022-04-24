package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) CreatePegawai(ctx context.Context, request pegawai_dto.CreatePegawaiRequestDto) (*pegawai_dto.CreatePegawaiResponseDto, error) {
	pegawai := createPegawaiFormRequest(request.Pegawai)
	err := s.RepositoryHolder.PegawaiRepository.Create(s.SharedHolder.Sql, &pegawai)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `Create`")
		return nil, err
	}
	return &pegawai_dto.CreatePegawaiResponseDto{
		MandatoryRequestV2Dto: request.MandatoryRequestV2Dto,
		Pegawai:               pegawai,
	}, err

}
