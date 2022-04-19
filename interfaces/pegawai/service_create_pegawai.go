package pegawai

import (
	"context"
	"github.com/pkg/errors"

	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) CreatePegawai(ctx context.Context, request pegawai_dto.CreatePegawaiRequestDto) (*pegawai_dto.CreatePegawaiResponseDto, error) {

	res, err := s.ApplicationHolder.PegawaiService.CreatePegawai(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed CreatePegawai")
	}

	return res, nil

}
