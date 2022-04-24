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

func Test_Service_GetAllPegawai(t *testing.T) {
	var (
		ctrl       = gomock.NewController(t)
		repo       = NewMockPegawaiRepository(ctrl)
		errMsg     = "Err"
		sqlMock    = mock_persistent.NewMockORM(ctrl)
		loggerMock = logger_mock.NewMockLogger(ctrl)

		pegawai []Pegawai
	)

	defer ctrl.Finish()

	t.Run("when GetAllPegawai failed", func(t *testing.T) {

		//repo.EXPECT().Find(gomock.Any()).Return(nil, errors.New(errMsg))
		repo.EXPECT().Find(sqlMock).Return(nil, errors.New(errMsg))
		loggerMock.EXPECT().Error("Error repository")

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{PegawaiRepository: repo})

		_, err := service.GetAllPegawai(context.Background(), pegawai_dto.GetAllPegawaiRequestDto{})
		assert.EqualError(t, err, errMsg)
	})

	t.Run("when GetAllPegawai success", func(t *testing.T) {
		repo.EXPECT().Find(sqlMock).Return(pegawai, nil)

		service, _ := NewService(shared.Holder{Sql: sqlMock}, RepositoryHolder{PegawaiRepository: repo})

		res, err := service.GetAllPegawai(context.Background(), pegawai_dto.GetAllPegawaiRequestDto{})
		assert.Equal(t, pegawai, res.Data)
		assert.Nil(t, err)
	})

}
