package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) CreatePegawai(ctx context.Context, request pegawai_dto.CreatePegawaiRequestDto) (*pegawai_dto.CreatePegawaiResponseDto, error) {
	pegawai := createPegawaiFormRequest(request.Pegawai)
	err := s.RepositoryHolder.PegawaiRepository.Create(s.SharedHolder.Sql, &pegawai)

	return nil, err

}
