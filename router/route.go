package router

import (
	"backend-b-payment-monitoring/config/db"
	"backend-b-payment-monitoring/controller"
	createPaymentController "backend-b-payment-monitoring/controller/payment/create-controller"
	"backend-b-payment-monitoring/repository"
	"backend-b-payment-monitoring/repository/loginRepo"
	createPayment "backend-b-payment-monitoring/repository/payment/create-repository"
	"backend-b-payment-monitoring/repository/paymentRepo"
	"backend-b-payment-monitoring/usecase"
	"backend-b-payment-monitoring/usecase/loginUsecase"
	createPaymentService "backend-b-payment-monitoring/usecase/payment/create-service"
	"backend-b-payment-monitoring/usecase/paymentUsecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New() // CALL LIBRARY GIN GONIC FOR ROUTER

	DB := db.DB // CALL FUNCTION DB
	fmt.Println(DB)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CALL DEPENDENCY REPOSITORY IN ABOVE
	roleRepo := repository.NewRoleRepoImpl(DB)
	accountRepo := repository.NewAccountRepoImpl(DB)
	loginRepo := loginRepo.NewLoginRepoImpl(DB)
	paymentRepo := paymentRepo.NewPaymentImpl(DB)
	createPaymentRepository := createPayment.NewRepositoryCreate(DB)


	// CALL DEPENDENCY USE CASE IN ABOVE
	roleUsecase := usecase.NewRoleUsecaseImpl(roleRepo)
	accountUsecase := usecase.NewAccountUsecaseImpl(accountRepo)
	loginUsecase := loginUsecase.NewLoginUsecaseImpl(loginRepo)
	paymentUsecase := paymentUsecase.NewRequestImpl(paymentRepo)
	createPaymentService := createPaymentService.NewPaymentCreate(createPaymentRepository)


	// create router group
	//newRoute := router.Group("/example")
	newRouteRole := router.Group("")

	// CALL DEPENDENCY CONTROLLER IN ABOVE
	controller.NewLoginControllerImpl(newRouteRole, loginUsecase)
	controller.NewRoleControllerImpl(newRouteRole, roleUsecase)
	controller.NewAccountControllerImpl(newRouteRole, accountUsecase)
	controller.NewPaymentControllerImpl(newRouteRole, paymentUsecase)
	createPaymentController.NewController(newRouteRole, createPaymentService)

	// RETURN ROUTER
	return router
}
