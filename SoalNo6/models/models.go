package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	KodeProduk string `json:"kodeProduk" gorm:"primaryKey"`
	NamaProduk string `json:"namaProduk" validate:"required"`
	Kuantitas  int    `json:"password" validate:"required"`
}

type Carts struct {
	gorm.Model
	Keranjang []Cart
}
