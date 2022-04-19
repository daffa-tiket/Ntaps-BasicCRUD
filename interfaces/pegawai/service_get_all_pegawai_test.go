package pegawai

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/pkg/errors"
	pegawai_application "github.com/tiket/learn-crud/application/pegawai"

	"github.com/tiket/learn-crud/application"
	"github.com/tiket/learn-crud/shared"

	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
)

func Test_GetAllPegawai(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to process GetAllPegawai", func(t *testing.T) {
		var (
			request = pegawai_dto.GetAllPegawaiRequestDto{}
			ctx     = context.Background()
		)

		as := pegawai_application.NewMockService(ctrl)
		as.EXPECT().GetAllPegawai(ctx, request).Return(nil, errors.New("failed"))

		service, _ := NewViewService(shared.Holder{}, application.Holder{
			PegawaiService: as,
		})
		_, err := service.GetAllPegawai(ctx, request)

		assert.EqualError(t, err, "failed to proceed GetAllPegawai: failed")

	})

	t.Run("when success to process GetAllPegawai", func(t *testing.T) {
		var (
			request = pegawai_dto.GetAllPegawaiRequestDto{}
			ctx     = context.Background()
		)

		as := pegawai_application.NewMockService(ctrl)
		as.EXPECT().GetAllPegawai(ctx, request).Return(&pegawai_dto.GetAllPegawaiResponseDto{}, nil)

		service, _ := NewViewService(shared.Holder{}, application.Holder{
			PegawaiService: as,
		})
		_, err := service.GetAllPegawai(ctx, request)

		assert.Nil(t, err)

	})
}
