package pegawai

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_GetAllPegawai(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when GetAllPegawai failed", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		_, err := service.GetAllPegawai(context.Background(), pegawai_dto.GetAllPegawaiRequestDto{})
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("when GetAllPegawai success", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		res, err := service.GetAllPegawai(context.Background(), pegawai_dto.GetAllPegawaiRequestDto{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

}
