package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) DeletePegawai(ctx context.Context, request pegawai_dto.DeletePegawaiRequestDto) (*pegawai_dto.DeletePegawaiResponseDto, error) {
	pegawai, err := s.RepositoryHolder.PegawaiRepository.FindOne(s.SharedHolder.Sql, "id = ?", request.ID)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `FindOne`")
		return nil, err
	}
	err = s.RepositoryHolder.PegawaiRepository.Delete(s.SharedHolder.Sql, pegawai)
	if err != nil {
		s.SharedHolder.Logger.Error("Error repository `Delete`")
		return nil, err
	}
	return nil, nil

}
