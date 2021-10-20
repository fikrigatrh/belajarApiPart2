package controller

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	role usecase.RoleUsecaseInterface
}

func NewRoleControllerImpl(r *gin.RouterGroup, role usecase.RoleUsecaseInterface) {
	handler := RoleController{role: role}

	r.GET("/getAllRole", handler.GetAllRole)
	r.POST("/role", handler.AddRoleController)
	r.GET("/role/:id", handler.GetRoleController)
	r.PUT("/role/:id", handler.UpdateRoleController)
	r.DELETE("/role/:id", handler.DeleteRoleController)
}

func (r RoleController) AddRoleController(c *gin.Context) {
	var role models.Role

	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}

	roleData, err := r.role.AddRoleUsecase(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseErrorCustom{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (r RoleController) GetRoleController(c *gin.Context) {
	id := c.Param("id")

	roleData := r.role.GetRoleUsecase(id)
	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    roleData,
	}
	c.JSON(http.StatusOK, response)
}

func (r RoleController) UpdateRoleController(c *gin.Context) {
	var role models.Role
	id := c.Param("id") // DARI ENDPOINT YANG ATAS DENGAN KATA KUNCI :ID

	err := c.ShouldBindJSON(&role) //fungsi untuk mengecek format JSON req. body, sudah sesuai
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return //end here. tidak dilanjut ke proses berikutnya
	}

	roleData, err := r.role.UpdateRoleUsecase(role, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when update")
		return
	}

	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    roleData,
	}
	c.JSON(http.StatusOK,  response)
}


func (r RoleController) GetAllRole(c *gin.Context) {
	getAllData, err := r.role.GetAllRoleUsecase()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error when binding")
		return
	}
	
	response := models.ResponseCustom{
		Status:  200,
		Message: "Berhasil",
		Data:    getAllData,
	}
	c.JSON(http.StatusOK, response)
}

func (r RoleController) DeleteRoleController(c *gin.Context) {
	var role models.Role
	id := c.Param("id")

	r.role.DeleteRoleUsecase(role, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "berhasil",
		"data":    "id_role " + id,
	})
}
