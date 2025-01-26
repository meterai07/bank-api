package repository

import (
	"bank/src/models"
	"errors"

	"gorm.io/gorm"
)

type NasabahRepository struct {
	db *gorm.DB
}

var (
	ErrDuplicateData    = errors.New("data sudah terdaftar")
	ErrAccountNotFound  = errors.New("rekening tidak ditemukan")
	ErrInsufficientFund = errors.New("saldo tidak mencukupi")
)

func NewNasabahRepository(db *gorm.DB) *NasabahRepository {
	return &NasabahRepository{db: db}
}

func (r *NasabahRepository) AutoMigrate() error {
	return r.db.AutoMigrate(&models.Nasabah{})
}

func (r *NasabahRepository) Create(nasabah *models.Nasabah) error {
	var count int64
	r.db.Model(&models.Nasabah{}).Where("nik = ? OR no_hp = ?", nasabah.NIK, nasabah.NoHP).Count(&count)
	if count > 0 {
		return ErrDuplicateData
	}

	return r.db.Create(nasabah).Error
}

func (r *NasabahRepository) Tabung(noRekening string, nominal int64) (int64, error) {
	var nasabah models.Nasabah
	tx := r.db.Begin()

	if err := tx.Where("no_rekening = ?", noRekening).First(&nasabah).Error; err != nil {
		tx.Rollback()
		return 0, ErrAccountNotFound
	}

	nasabah.Saldo += nominal
	if err := tx.Save(&nasabah).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return nasabah.Saldo, nil
}

func (r *NasabahRepository) Tarik(noRekening string, nominal int64) (int64, error) {
	var nasabah models.Nasabah
	tx := r.db.Begin()

	if err := tx.Where("no_rekening = ?", noRekening).First(&nasabah).Error; err != nil {
		tx.Rollback()
		return 0, ErrAccountNotFound
	}

	if nasabah.Saldo < nominal {
		tx.Rollback()
		return 0, ErrInsufficientFund
	}

	nasabah.Saldo -= nominal
	if err := tx.Save(&nasabah).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return nasabah.Saldo, nil
}

func (r *NasabahRepository) GetSaldo(noRekening string) (int64, error) {
	var nasabah models.Nasabah
	if err := r.db.Select("saldo").Where("no_rekening = ?", noRekening).First(&nasabah).Error; err != nil {
		return 0, ErrAccountNotFound
	}
	return nasabah.Saldo, nil
}

func (r *NasabahRepository) FindByNIK(nik string) (*models.Nasabah, error) {
	var nasabah models.Nasabah
	if err := r.db.Where("nik = ?", nik).First(&nasabah).Error; err != nil {
		return nil, ErrAccountNotFound
	}
	return &nasabah, nil
}
