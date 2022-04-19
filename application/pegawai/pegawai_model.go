package pegawai

import (
	"time"
)

type (
	Pegawai struct {
		ID          uint      `gorm:"primary_key"`
		CreatedDate time.Time `gorm:"column:created_date"`
		UpdatedDate time.Time `gorm:"column:updated_date"`
		CreatedBy   string    `gorm:"column:created_by"`
		UpdatedBy   string    `gorm:"column:updated_by"`
		IsDeleted   bool      `gorm:"column:is_deleted"`
		Version     int       `gorm:"column:version"`
		Nama        string    `gorm:"column:nama"`
		Alamat      string    `gorm:"column:alamat"`
		Telepon     string    `gorm:"column:telepon"`
	}
)

func (m *Pegawai) TableName() string {
	return "pegawai"
}
