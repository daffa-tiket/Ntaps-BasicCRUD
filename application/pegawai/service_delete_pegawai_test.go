package pegawai

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_persistent "github.com/tiket/TIX-HOTEL-UTILITIES-GO/mocks/persistent"
	logger_mock "github.com/tiket/learn-crud/mocks/logger"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_DeletePegawai(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		repo        = NewMockPegawaiRepository(ctrl)
		errMsg      = "Err"
		sqlMock     = mock_persistent.NewMockORM(ctrl)
		loggerMock  = logger_mock.NewMockLogger(ctrl)
		pegawaiMock = Pegawai{}
	)

	defer ctrl.Finish()

	t.Run("when FindOnePegawai failed", func(t *testing.T) {
		repo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(nil, errors.New(errMsg))
		loggerMock.EXPECT().Error("Error repository `FindOne`")

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{
			PegawaiRepository: repo,
		})

		_, err := service.UpdatePegawai(context.Background(), pegawai_dto.UpdatePegawaiRequestDto{
			ID: 1,
		})
		assert.EqualError(t, err, errMsg)
	})

	t.Run("when DeletePegawai failed", func(t *testing.T) {
		repo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(&pegawaiMock, nil)
		repo.EXPECT().Delete(sqlMock, &pegawaiMock).Return(errors.New(errMsg))
		loggerMock.EXPECT().Error("Error repository `Delete`")

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{
			PegawaiRepository: repo,
		})

		res, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{ID: 1})
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

	t.Run("when DeletePegawai success", func(t *testing.T) {
		repo.EXPECT().FindOne(sqlMock, "id = ?", 1).Return(&pegawaiMock, nil)
		repo.EXPECT().Delete(sqlMock, &pegawaiMock).Return(nil)

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{
			PegawaiRepository: repo,
		})

		res, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{ID: 1})
		assert.Nil(t, res)
		assert.Nil(t, err)
	})

}
