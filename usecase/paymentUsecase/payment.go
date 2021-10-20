package paymentUsecase

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/repository"
	"backend-b-payment-monitoring/usecase"
)

type PaymentUsecaseStruct struct {
	RequestRepo repository.PaymentRepoInterface
}

func NewRequestImpl(RequestRepo repository.PaymentRepoInterface) usecase.PaymentUsecaseInterface {
	return &PaymentUsecaseStruct{RequestRepo}
}

func (ru PaymentUsecaseStruct) GetAllRequestUsecase() ([]models.EntityPayment, error) {
	getAllReq, err := ru.RequestRepo.GetAllRequest()
	if err != nil {
		return nil, err
	}

	return getAllReq, nil
}

func (ru PaymentUsecaseStruct) GetReqByunit(id string) models.EntityPayment {
	ReqData := ru.RequestRepo.GetReqByunit(id)

	return ReqData
}

func (ru PaymentUsecaseStruct) UpdateAprovalUsecase(aproval models.EntityPayment, tipeReqRes string) (models.EntityPayment, error) {

	payment, err := ru.RequestRepo.GetReqByIdPayment(int(aproval.IdPayment))
	if err != nil {
		return models.EntityPayment{}, err
	}

	aproval.IdPayment = payment.IdPayment
	aproval.IdUnit = payment.IdUnit
	aproval.DimintaOleh = payment.DimintaOleh
	aproval.Keperluan = payment.Keperluan
	aproval.TanggalPembayaranAktual = payment.TanggalPembayaranAktual
	aproval.JumlahPayment = payment.JumlahPayment
	aproval.Terbilang = payment.Terbilang
	aproval.NamaRekPenerima = payment.NamaRekPenerima
	aproval.NoRekPenerima = payment.NoRekPenerima
	aproval.RequestTerkirim = payment.RequestTerkirim

	if tipeReqRes == "direject_gs" {
		aproval.StatusRequest = 2
	} else if tipeReqRes  == "diteruskan" {
		aproval.StatusRequest = 3
	} else if tipeReqRes  == "direject_ac" {
		aproval.StatusRequest = 4
	} else if tipeReqRes  == "disetujui" {
		aproval.StatusRequest = 5
	}

	//TODO
	repo, err := ru.RequestRepo.UpdateAprovalRepo(aproval)
	if err != nil {
		return models.EntityPayment{}, err
	}

	return repo, nil
}

func (ru PaymentUsecaseStruct) AddStatusReqUsecase(status models.StatusRequest) models.StatusRequest {
	statusReq := ru.RequestRepo.AddStatusReqRepo(status)

	return statusReq
}
