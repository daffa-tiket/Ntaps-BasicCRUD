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

func Test_CreatePegawai(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to process CreatePegawai", func(t *testing.T) {
		var (
			request = pegawai_dto.CreatePegawaiRequestDto{}
			ctx     = context.Background()
		)

		as := pegawai_application.NewMockService(ctrl)
		as.EXPECT().CreatePegawai(ctx, request).Return(nil, errors.New("failed"))

		service, _ := NewViewService(shared.Holder{}, application.Holder{
			PegawaiService: as,
		})
		_, err := service.CreatePegawai(ctx, request)

		assert.EqualError(t, err, "failed to proceed CreatePegawai: failed")

	})

	t.Run("when success to process CreatePegawai", func(t *testing.T) {
		var (
			request = pegawai_dto.CreatePegawaiRequestDto{}
			ctx     = context.Background()
		)

		as := pegawai_application.NewMockService(ctrl)
		as.EXPECT().CreatePegawai(ctx, request).Return(&pegawai_dto.CreatePegawaiResponseDto{}, nil)

		service, _ := NewViewService(shared.Holder{}, application.Holder{
			PegawaiService: as,
		})
		_, err := service.CreatePegawai(ctx, request)

		assert.Nil(t, err)

	})
}
