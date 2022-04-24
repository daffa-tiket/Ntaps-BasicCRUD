package pegawai

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_persistent "github.com/tiket/TIX-HOTEL-UTILITIES-GO/mocks/persistent"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_DeletePegawai(t *testing.T) {
	var (
		ctrl    = gomock.NewController(t)
		sqlMock = mock_persistent.NewMockORM(ctrl)
	)

	defer ctrl.Finish()

	t.Run("when FindOne DeetePegawai failed", func(t *testing.T) {

		pegawaiRepo := NewMockPegawaiRepository(ctrl)
		pegawaiRepo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(nil, errors.New("Error FindOne"))

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{
			PegawaiRepository: pegawaiRepo,
		})

		_, err := service.UpdatePegawai(context.Background(), pegawai_dto.UpdatePegawaiRequestDto{
			ID: 1,
		})
		assert.EqualError(t, err, "Error FindOne")
	})

	t.Run("when DeletePegawai failed", func(t *testing.T) {
		pegawaiMock := Pegawai{
			Nama:    "Nando",
			Alamat:  "Bandung",
			Telepon: "0813",
		}
		pegawaiRepo := NewMockPegawaiRepository(ctrl)
		pegawaiRepo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(&pegawaiMock, nil)
		pegawaiRepo.EXPECT().Delete(sqlMock, &pegawaiMock).Return(errors.New("Error Delete"))

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{
			PegawaiRepository: pegawaiRepo,
		})

		res, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{ID: 1})
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("when DeletePegawai success", func(t *testing.T) {
		pegawaiMock := Pegawai{
			Nama:    "Nando",
			Alamat:  "Bandung",
			Telepon: "0813",
		}
		pegawaiRepo := NewMockPegawaiRepository(ctrl)
		pegawaiRepo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(&pegawaiMock, nil)
		pegawaiRepo.EXPECT().Delete(sqlMock, &pegawaiMock).Return(nil)

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{
			PegawaiRepository: pegawaiRepo,
		})

		res, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{ID: 1})
		assert.Nil(t, res)
		assert.Nil(t, err)
	})

}
