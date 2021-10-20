package controller

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type PaymentController struct {
	Payment usecase.PaymentUsecaseInterface
}

func NewPaymentControllerImpl(p *gin.RouterGroup, Payment usecase.PaymentUsecaseInterface) {
	handler := PaymentController{Payment}

	p.GET("/getAllRequestPayment", handler.GetAllRequest)
	p.GET("/list_payment", handler.GetReqByUnit)
	p.POST("/statusPayment", handler.AddStatusReqController)
	p.POST("/payment/status/:type", handler.ApprovalReqPayment)
}

func (pc PaymentController) GetAllRequest(c *gin.Context) {
	getAllRequest, err := pc.Payment.GetAllRequestUsecase()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when get all request Payment")
		return

	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    getAllRequest,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PaymentController) GetReqByUnit(c *gin.Context)  {
	id := c.Query("id")

	reqData := pc.Payment.GetReqByunit(id)
	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    reqData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PaymentController) AddStatusReqController(c *gin.Context) {
	var reqStatus models.StatusRequest

	err := c.ShouldBindJSON(&reqStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}


	log.Println(reqStatus.StatusName, "<<< statusReq")
	statusReqData := pc.Payment.AddStatusReqUsecase(reqStatus)
	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    statusReqData,
	}
	c.JSON(http.StatusOK, response)
}

func (pc PaymentController) ApprovalReqPayment(c *gin.Context) {
	var request models.EntityPayment

	err := c.ShouldBindJSON(&request)
	if err != nil {
		return
	}

	tipeReqTemp := c.Param("type")

	tipeReqRes := strings.ToLower(tipeReqTemp)

	approval, err := pc.Payment.UpdateAprovalUsecase(request, tipeReqRes)
	if err != nil {
		return
	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    approval,
	}
	c.JSON(http.StatusOK, response)

}