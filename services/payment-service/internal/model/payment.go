package model

import (
	"gorm.io/gorm"
)

type PaymentTransaction struct {
	gorm.Model
	TxnRef    string `gorm:"type:varchar(100);unique" json:"txn_ref"` // unique
	BookingID uint
	Status    string // PENDING, SUCCESS, FAILED
}
