package pegawai

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tiket/learn-crud/shared"
	"testing"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/persistent"

	mock_orm "github.com/tiket/learn-crud/mocks/postgres"
)

func Test_Pegawai_WithTransaction(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to rollback", func(t *testing.T) {
		var (
			orm   = mock_orm.NewMockORM(ctrl)
			block = func(tx persistent.ORM) error {
				return errors.New("callback failed")
			}
		)

		orm.EXPECT().Begin().Return(orm)
		orm.EXPECT().Rollback().Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{Sql: orm})
		err := repo.WithTransaction(block)

		assert.EqualError(t, err, "failed to rollback: failed")
	})

	t.Run("when success to rollback", func(t *testing.T) {
		var (
			orm   = mock_orm.NewMockORM(ctrl)
			block = func(tx persistent.ORM) error {
				return errors.New("callback failed")
			}
		)

		orm.EXPECT().Begin().Return(orm)
		orm.EXPECT().Rollback().Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{Sql: orm})
		err := repo.WithTransaction(block)

		assert.EqualError(t, err, "block failed: callback failed")
	})

	t.Run("when failed to commit", func(t *testing.T) {
		var (
			orm   = mock_orm.NewMockORM(ctrl)
			block = func(tx persistent.ORM) error {
				return nil
			}
		)

		orm.EXPECT().Begin().Return(orm)
		orm.EXPECT().Commit().Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{Sql: orm})
		err := repo.WithTransaction(block)

		assert.EqualError(t, err, "failed to commit: failed")
	})

	t.Run("when success to commit", func(t *testing.T) {
		var (
			orm   = mock_orm.NewMockORM(ctrl)
			block = func(tx persistent.ORM) error {
				return nil
			}
		)

		orm.EXPECT().Begin().Return(orm)
		orm.EXPECT().Commit().Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{Sql: orm})
		err := repo.WithTransaction(block)

		assert.Nil(t, err)
	})
}

func Test_Pegawai_Find(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to find", func(t *testing.T) {
		var (
			orm = mock_orm.NewMockORM(ctrl)
		)

		orm.EXPECT().All(gomock.Any()).Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{})
		res, err := repo.Find(orm)

		assert.EqualError(t, err, "failed to find pegawai: failed")
		assert.Nil(t, res)
	})

	t.Run("when success to find", func(t *testing.T) {
		var (
			orm = mock_orm.NewMockORM(ctrl)
		)

		orm.EXPECT().All(gomock.Any()).Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{})
		res, err := repo.Find(orm)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func Test_Pegawai_FindOne(t *testing.T) {
	var (
		ctrl  = gomock.NewController(t)
		query interface{}
		args  interface{}
	)

	defer ctrl.Finish()

	t.Run("when failed to Find", func(t *testing.T) {
		var (
			orm = mock_orm.NewMockORM(ctrl)
		)

		orm.EXPECT().Where(query, args).Return(orm)
		orm.EXPECT().First(gomock.Any()).Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{})
		_, err := repo.FindOne(orm, query, args)

		assert.EqualError(t, err, "failed to find one pegawai: failed")
	})

	t.Run("when success to Find", func(t *testing.T) {
		var (
			orm = mock_orm.NewMockORM(ctrl)
		)

		orm.EXPECT().Where(query, args).Return(orm)
		orm.EXPECT().First(gomock.Any()).Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{})
		_, err := repo.FindOne(orm, query, args)

		assert.Nil(t, err)
	})
}

func Test_Pegawai_Create(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to Create", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Create(&object).Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Create(orm, &object)

		assert.EqualError(t, err, "failed to create pegawai: failed")
	})

	t.Run("when success to Create", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Create(&object).Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Create(orm, &object)

		assert.Nil(t, err)
	})
}

func Test_Pegawai_Update(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to Update", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Update(&object).Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Update(orm, &object)

		assert.EqualError(t, err, "failed to update pegawai: failed")
	})

	t.Run("when success to Update", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Update(&object).Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Update(orm, &object)

		assert.Nil(t, err)
	})
}

func Test_Pegawai_Delete(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("when failed to Delete", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Delete(&object).Return(errors.New("failed"))

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Delete(orm, &object)

		assert.EqualError(t, err, "failed to delete pegawai: failed")
	})

	t.Run("when success to Delete", func(t *testing.T) {
		var (
			orm    = mock_orm.NewMockORM(ctrl)
			object = Pegawai{}
		)

		orm.EXPECT().Delete(&object).Return(nil)

		repo, _ := NewPegawaiRepository(shared.Holder{})
		err := repo.Delete(orm, &object)

		assert.Nil(t, err)
	})
}
