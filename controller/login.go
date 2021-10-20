package controller

import (
	"backend-b-payment-monitoring/auth"
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	LoginUsecase usecase.LoginUsecaseInterface
}

func NewLoginControllerImpl(r *gin.RouterGroup, LoginUsecase usecase.LoginUsecaseInterface) {
	handler := LoginController{LoginUsecase}
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
}

func (l LoginController) Login(c *gin.Context) {
	var user models.OfficerAccount

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var authD models.Auth

	authRes, err := l.LoginUsecase.CreateAuth(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	authD.AuthUUID = authRes.AuthUUID
	authD.Username = authRes.Username
	authD.RoleId = authRes.RoleId
	authD.RoleName = authRes.RoleName

	var jwt models.JwtModel

	in, err := l.LoginUsecase.SignIn(authD)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	jwt.IdOffice = authRes.IdOffice
	jwt.IdCust = authRes.IdCust
	jwt.Token = in

	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    jwt,
	}
	c.JSON(http.StatusOK, response)

}

func (l LoginController) Logout(c *gin.Context) {
	tokenAuth, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	err = l.LoginUsecase.DeleteAuth(tokenAuth.AuthUUID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    tokenAuth,
	}
	c.JSON(http.StatusOK, response)
}