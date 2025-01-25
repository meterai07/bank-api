package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Nasabah struct {
	gorm.Model
	NoRekening string `gorm:"uniqueIndex;size:40;not null" json:"no_rekening"`
	Nama       string `gorm:"size:255;not null" json:"nama"`
	NIK        string `gorm:"uniqueIndex;size:16;not null" json:"nik"`
	NoHP       string `gorm:"uniqueIndex;size:15;not null" json:"no_hp"`
	Saldo      int64  `gorm:"not null;default:0" json:"saldo"`
}

func (n *Nasabah) BeforeCreate(tx *gorm.DB) error {
	n.NoRekening = "REK-" + uuid.New().String()
	return nil
}
