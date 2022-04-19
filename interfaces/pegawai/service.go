// Code generated by ntaps. DO NOT EDIT.
package pegawai

import (
	"context"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"

	"github.com/tiket/learn-crud/application"
	"github.com/tiket/learn-crud/shared"
	// - pegawai-import-end
)

type (
	ViewService interface {

		// - pegawai-get-all-pegawai-start
		GetAllPegawai(ctx context.Context, request pegawai_dto.GetAllPegawaiRequestDto) (*pegawai_dto.GetAllPegawaiResponseDto, error)
		// - pegawai-get-all-pegawai-end

		// - pegawai-update-pegawai-start
		UpdatePegawai(ctx context.Context, request pegawai_dto.UpdatePegawaiRequestDto) (*pegawai_dto.UpdatePegawaiResponseDto, error)
		// - pegawai-update-pegawai-end

		// - pegawai-create-pegawai-start
		CreatePegawai(ctx context.Context, request pegawai_dto.CreatePegawaiRequestDto) (*pegawai_dto.CreatePegawaiResponseDto, error)
		// - pegawai-create-pegawai-end

		// - pegawai-delete-pegawai-start
		DeletePegawai(ctx context.Context, request pegawai_dto.DeletePegawaiRequestDto) (*pegawai_dto.DeletePegawaiResponseDto, error)
		// - pegawai-delete-pegawai-end

		// - pegawai-view-service-end
	}

	service struct {
		SharedHolder      shared.Holder
		ApplicationHolder application.Holder
	}
)

func NewViewService(sh shared.Holder, ah application.Holder) (ViewService, error) {
	return &service{sh, ah}, nil
}
