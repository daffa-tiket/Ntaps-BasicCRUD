package pegawai

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tiket/learn-crud/shared"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"testing"
)

func Test_Service_DeletePegawai(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when DeletePegawai failed", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		_, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{})
		assert.EqualError(t, err, "not implemented")
	})

	t.Run("when DeletePegawai success", func(t *testing.T) {

		service, _ := NewService(shared.Holder{}, RepositoryHolder{})

		res, err := service.DeletePegawai(context.Background(), pegawai_dto.DeletePegawaiRequestDto{})
		assert.Nil(t, res)
		assert.NotNil(t, err)
	})

}
