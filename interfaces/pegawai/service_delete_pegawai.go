package pegawai

import (
	"context"
	"github.com/pkg/errors"

	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func (s *service) DeletePegawai(ctx context.Context, request pegawai_dto.DeletePegawaiRequestDto) (*pegawai_dto.DeletePegawaiResponseDto, error) {

	res, err := s.ApplicationHolder.PegawaiService.DeletePegawai(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed DeletePegawai")
	}

	return res, nil

}
