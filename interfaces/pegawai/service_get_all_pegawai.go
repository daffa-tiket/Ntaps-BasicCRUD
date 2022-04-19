package pegawai

import (
	"context"
	"github.com/pkg/errors"

	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) GetAllPegawai(ctx context.Context, request pegawai_dto.GetAllPegawaiRequestDto) (*pegawai_dto.GetAllPegawaiResponseDto, error) {

	res, err := s.ApplicationHolder.PegawaiService.GetAllPegawai(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed GetAllPegawai")
	}

	return res, nil

}
