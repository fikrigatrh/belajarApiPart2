package models

import (
	"gorm.io/gorm"
	"time"
)

type Approval struct {
	gorm.Model
	IdPayment     int       `json:"id_payment"`
	StatusRequest int       `json:"status_request"`
	Reason        string    `json:"reason"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type StatusRequest struct {
	Status     int    `json:"status"`
	StatusName string `json:"status_name"`
}

type ResponseCustom struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorCustom struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type EntityPayment struct {
	IdPayment               uint      `gorm:"primarykey;auto_increment;not null" json:"id_payment"`
	IdUnit                  uint      `gorm:"type:int(10);auto_increment; not null;" json:"id_unit"`
	DimintaOleh             string    `gorm:"type:varchar(255);not null;" json:"diminta_oleh"`
	Keperluan               string    `gorm:"type:varchar(255);not null;" json:"keperluan"`
	TanggalPembayaranAktual string    `json:"tanggal_pembayaran_aktual"`
	JumlahPayment           int64     `gorm:"not null;" json:"jumlah_payment"`
	Terbilang               string    `gorm:"type:varchar(255);not null;" json:"terbilang"`
	NamaRekPenerima         string    `gorm:"type:varchar(255);not null;" json:"nama_rek_penerima"`
	NoRekPenerima           string    `gorm:"type:varchar(255);not null;" json:"no_rek_penerima"`
	CreatedAt               time.Time `json:"-"`
	UpdatedAt               time.Time `json:"-"`
	StatusRequest           int       `json:"status"`
	Reason                  string    `json:"reason,omitempty"`
	RequestTerkirim         string    `json:"request_terkirim,omitempty"`
}
