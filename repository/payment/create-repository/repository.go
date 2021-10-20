package createPaymentRepository

import (
	"backend-b-payment-monitoring/models"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	CreatePaymentRepository(input *models.EntityPayment) (*models.EntityPayment, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreatePaymentRepository(input *models.EntityPayment) (*models.EntityPayment, string) {
	var payment models.EntityPayment
	db := r.db.Model(&payment)
	errorCode := make(chan string, 1)

	layout := "2006-1-02 15:04:05"
	tm := time.Now()
	res := tm.Format(layout)

	payment.IdUnit = input.IdUnit
	payment.DimintaOleh = input.DimintaOleh
	payment.Keperluan = input.Keperluan
	payment.TanggalPembayaranAktual = input.TanggalPembayaranAktual
	payment.JumlahPayment = input.JumlahPayment
	payment.Terbilang = input.Terbilang
	payment.NamaRekPenerima = input.NamaRekPenerima
	payment.NoRekPenerima = input.NoRekPenerima
	payment.StatusRequest = 1
	payment.RequestTerkirim = res

	createNewPayment := db.Debug().Create(&payment)
	db.Commit()

	if createNewPayment.Error != nil {
		errorCode <- "CREATE_PAYMENT_FAILED"
		return &payment, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &payment, <-errorCode
}
