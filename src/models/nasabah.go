package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Nasabah struct {
	gorm.Model
	NoRekening string `gorm:"uniqueIndex;size:40;not null" json:"no_rekening"`
	Nama       string `gorm:"size:255;not null" json:"nama"`
	NIK        string `gorm:"uniqueIndex;size:16;not null" json:"nik"`
	NoHP       string `gorm:"uniqueIndex;size:15;not null" json:"no_hp"`
	Password   string `gorm:"size:255;not null" json:"-"`
	Saldo      int64  `gorm:"not null;default:0" json:"saldo"`
}

func (n *Nasabah) BeforeCreate(tx *gorm.DB) error {
	n.NoRekening = "REK-" + uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	n.Password = string(hashedPassword)

	return nil
}
