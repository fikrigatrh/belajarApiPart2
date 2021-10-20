package createPaymentService

import (
	"backend-b-payment-monitoring/models"
	createPayment "backend-b-payment-monitoring/repository/payment/create-repository"
	"time"
)

type Service interface {
	CreatePaymentService(input *createPayment.InputCreatePayment) (*models.EntityPayment, string)
}

type service struct {
	repository createPayment.Repository
}

func NewPaymentCreate(repository createPayment.Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreatePaymentService(input *createPayment.InputCreatePayment) (*models.EntityPayment, string) {

	layout := "2006-1-02 15:04:05"
	tm := time.Now()
	res := tm.Format(layout)

	payments := models.EntityPayment{
		IdUnit:                  input.IdUnit,
		DimintaOleh:             input.DimintaOleh,
		Keperluan:               input.Keperluan,
		TanggalPembayaranAktual: input.TanggalPembayaranAktual,
		JumlahPayment:           input.JumlahPayment,
		Terbilang:               input.Terbilang,
		NamaRekPenerima:         input.NamaRekPenerima,
		NoRekPenerima:           input.NoRekPenerima,
		StatusRequest:           1,
	}

	resultCreatePayment, errCreatePayment := s.repository.CreatePaymentRepository(&payments)
	resultCreatePayment.RequestTerkirim = res

	return resultCreatePayment, errCreatePayment
}
