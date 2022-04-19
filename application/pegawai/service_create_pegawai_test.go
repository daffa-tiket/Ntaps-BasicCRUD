package pegawai

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_CreatePegawai(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when CreatePegawai failed", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		_, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{})
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("when CreatePegawai success", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		res, err := service.CreatePegawai(context.Background(), pegawai_dto.CreatePegawaiRequestDto{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

}
