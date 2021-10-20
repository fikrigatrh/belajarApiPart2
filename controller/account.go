package controller

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	account usecase.AccountUsecaseInterface
}

func NewAccountControllerImpl(r *gin.RouterGroup, account usecase.AccountUsecaseInterface) {
	handler := AccountController{account: account}

	r.POST("/akun/create", handler.AddAccountController)
	r.GET("/list_akun", handler.GetAllAccountController)
	r.GET("/akun/:id", handler.GetAccountController)
	r.PUT("/akun/:id", handler.UpdateAccountController)
	r.DELETE("/akun/:id", handler.DeleteAccountController)
}

func (r AccountController) AddAccountController(c *gin.Context) {
	var accOfficer models.OfficerAccount

	if errOfficer := c.ShouldBindJSON(&accOfficer); errOfficer != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}

	accDataOfficer, accDataCustomer, err := r.account.AddAccountUsecase(accOfficer)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if accDataOfficer.Role == 0 {
		response := models.ResponseCustom{
			Status:  200,
			Message: "Berhasil",
			Data:    accDataCustomer,
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := models.ResponseCustom{
			Status:  200,
			Message: "Berhasil",
			Data:    accDataOfficer,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (r AccountController) GetAllAccountController(c *gin.Context) {
	getAllData, err := r.account.GetAllAccountUsecase()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}

	c.JSON(http.StatusOK, getAllData)
}

func (r AccountController) GetAccountController(c *gin.Context) {
	var account models.Role
	id := c.Param("id")

	accountData := r.account.GetAccountUsecase(account, id)
	c.JSON(http.StatusOK, accountData)
}

func (r AccountController) UpdateAccountController(c *gin.Context) {
	var account models.Role
	id := c.Param("id")

	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}

	accountData, err := r.account.UpdateAccountUsecase(account, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when update")
		return
	}
	c.JSON(http.StatusOK, accountData)
}

func (r AccountController) DeleteAccountController(c *gin.Context) {
	var account models.Role
	id := c.Param("id")

	r.account.DeleteAccountUsecase(account, id)
	c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
}
