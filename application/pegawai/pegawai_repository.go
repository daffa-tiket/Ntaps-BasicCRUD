package pegawai

import (
	"github.com/pkg/errors"
	"github.com/tiket/learn-crud/shared"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/persistent"
)

type (
	PegawaiRepository interface {
		Find(tx persistent.ORM) ([]Pegawai, error)
		FindOne(tx persistent.ORM, query interface{}, args ...interface{}) (*Pegawai, error)
		Create(tx persistent.ORM, object *Pegawai) error
		Update(tx persistent.ORM, object *Pegawai) error
		Delete(tx persistent.ORM, object *Pegawai) error
		WithTransaction(block func(tx persistent.ORM) error) error
		Tx() persistent.ORM
	}

	pegawai_repo struct {
		SharedHolder shared.Holder
	}
)

func (i *pegawai_repo) TableName() string {
	return "pegawai"
}

func (i *pegawai_repo) WithTransaction(block func(tx persistent.ORM) error) error {
	// - begin transaction
	trx := i.SharedHolder.Sql.Begin()

	if outer := block(trx); outer != nil {
		// - rollback
		if err := trx.Rollback(); err != nil {
			return errors.Wrap(err, "failed to rollback")
		}
		return errors.Wrap(outer, "block failed")
	} else {
		// - commit
		if err := trx.Commit(); err != nil {
			return errors.Wrap(err, "failed to commit")
		}
		return nil
	}
}

func (i *pegawai_repo) Tx() persistent.ORM {
	return i.SharedHolder.Sql.Begin()
}

func (i *pegawai_repo) Find(tx persistent.ORM) ([]Pegawai, error) {
	var (
		objects = make([]Pegawai, 0)
	)

	if err := tx.All(&objects); err != nil {
		return nil, errors.Wrapf(err, "failed to find %s", i.TableName())
	}

	return objects, nil
}

func (i *pegawai_repo) FindOne(tx persistent.ORM, query interface{}, args ...interface{}) (*Pegawai, error) {
	var (
		object = &Pegawai{}
	)
	if err := tx.Where(query, args...).First(object); err != nil {
		return nil, errors.Wrapf(err, "failed to find one %s", i.TableName())
	}

	return object, nil
}

func (i *pegawai_repo) Create(tx persistent.ORM, object *Pegawai) error {
	if err := tx.Create(object); err != nil {
		return errors.Wrapf(err, "failed to create %s", i.TableName())
	}
	return nil
}

func (i *pegawai_repo) Update(tx persistent.ORM, object *Pegawai) error {
	if err := tx.Update(object); err != nil {
		return errors.Wrapf(err, "failed to update %s", i.TableName())
	}
	return nil
}

func (i *pegawai_repo) Delete(tx persistent.ORM, object *Pegawai) error {
	if err := tx.Delete(object); err != nil {
		return errors.Wrapf(err, "failed to delete %s", i.TableName())
	}
	return nil
}

func NewPegawaiRepository(sh shared.Holder) (PegawaiRepository, error) {
	return &pegawai_repo{sh}, nil
}
