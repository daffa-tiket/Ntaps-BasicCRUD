package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) DeletePegawai(ctx context.Context, request pegawai_dto.DeletePegawaiRequestDto) (*pegawai_dto.DeletePegawaiResponseDto, error) {
	pegawai, _ := s.RepositoryHolder.PegawaiRepository.FindOne(s.SharedHolder.Sql, "id = ?", request.ID)
	err := s.RepositoryHolder.PegawaiRepository.Delete(s.SharedHolder.Sql, pegawai)
	return nil, err

}
