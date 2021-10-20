package createPaymentController

import (
	createPayment "backend-b-payment-monitoring/repository/payment/create-repository"
	createPaymentService "backend-b-payment-monitoring/usecase/payment/create-service"
	"backend-b-payment-monitoring/utils"
	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"net/http"
)

type controller struct {
	service createPaymentService.Service
}

func NewController(r *gin.RouterGroup, service createPaymentService.Service) {
	handler :=  controller{service: service}

	r.POST("payment/create", handler.CreatePaymentController)
}

func (c controller) CreatePaymentController(ctx *gin.Context) {
	var input createPayment.InputCreatePayment
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "oopsss")
		return
	}

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "IdUnit",
				Message: "id unit is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "IdUnit",
				Message: "id unit must be numeric",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "DimintaOleh",
				Message: "name is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "uppercase",
				Field:   "DimintaOleh",
				Message: "Diminta oleh must be using uppercase",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Keperluan",
				Message: "keperluan is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "TanggalPembayaranAktual",
				Message: "Tanggal pembayaran aktual is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "JumlahPayment",
				Message: "Jumlah payment is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "JumlahPayment",
				Message: "Jumlah payment must be using number",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Terbilang",
				Message: "terbilang  is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "NamaRekPenerima",
				Message: "Nama rekening penerima  is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "NoRekPenerima",
				Message: "No rekening penerima is required on body",
			},
			gpc.ErrorMetaConfig{
				Tag:     "numeric",
				Field:   "NoRekPenerima",
				Message: "No rek penerima is required on body",
			},
		},
	}

	errResponse, errCount := utils.GoValidator(&input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	result, errCreatePayment := c.service.CreatePaymentService(&input)

	switch errCreatePayment {

	case "CREATE_PAYMENT_FAILED":
		utils.APIResponse(ctx, "create new payment failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		utils.APIResponse(ctx, "create payment successfully", http.StatusOK, http.MethodPost, result)
	}


}
