package pegawai

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_persistent "github.com/tiket/TIX-HOTEL-UTILITIES-GO/mocks/persistent"
	logger_mock "github.com/tiket/learn-crud/mocks/logger"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_CreatePegawai(t *testing.T) {
	var (
		ctrl       = gomock.NewController(t)
		repo       = NewMockPegawaiRepository(ctrl)
		errMsg     = "Err"
		sqlMock    = mock_persistent.NewMockORM(ctrl)
		loggerMock = logger_mock.NewMockLogger(ctrl)
	)

	defer ctrl.Finish()

	t.Run("when CreatePegawai failed", func(t *testing.T) {
		repo.EXPECT().Create(sqlMock, gomock.Any()).Return(errors.New(errMsg))
		loggerMock.EXPECT().Error("Error repository `Create`")

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{
			PegawaiRepository: repo,
		})

		_, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{Pegawai: pegawai_dto.Pegawai{}})
		assert.EqualError(t, err, errMsg)
	})

	t.Run("when CreatePegawai success", func(t *testing.T) {
		repo.EXPECT().Create(sqlMock, gomock.Any()).Return(nil)

		service, _ := NewService(shared.Holder{Sql: sqlMock, Logger: loggerMock}, RepositoryHolder{
			PegawaiRepository: repo,
		})

		res, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{Pegawai: pegawai_dto.Pegawai{}})
		fmt.Println(res)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

}
