package pegawai

import (
	"context"
	"github.com/pkg/errors"

	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) UpdatePegawai(ctx context.Context, request pegawai_dto.UpdatePegawaiRequestDto) (*pegawai_dto.UpdatePegawaiResponseDto, error) {

	res, err := s.ApplicationHolder.PegawaiService.UpdatePegawai(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed UpdatePegawai")
	}

	return res, nil

}
