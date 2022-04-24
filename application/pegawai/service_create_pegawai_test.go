package pegawai

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_persistent "github.com/tiket/TIX-HOTEL-UTILITIES-GO/mocks/persistent"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_CreatePegawai(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		sqlMock = mock_persistent.NewMockORM(ctrl)
	)

	defer ctrl.Finish()

	t.Run("when CreatePegawai failed", func(t *testing.T) {
		pegawaiRepo := NewMockPegawaiRepository(ctrl)
		pegawaiRepo.EXPECT().Create(sqlMock, gomock.Any()).Return(errors.New("Error repository"))

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{
			PegawaiRepository: pegawaiRepo,
		})

		_, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{Pegawai: pegawai_dto.Pegawai{}})
		assert.EqualError(t, err, "Error repository")
	})

	t.Run("when CreatePegawai success", func(t *testing.T) {
		pegawaiRepo := NewMockPegawaiRepository(ctrl)
		pegawaiRepo.EXPECT().Create(sqlMock, gomock.Any()).Return(nil)

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{
			PegawaiRepository: pegawaiRepo,
		})

		res, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{Pegawai: pegawai_dto.Pegawai{}})
		fmt.Println(res)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

}
