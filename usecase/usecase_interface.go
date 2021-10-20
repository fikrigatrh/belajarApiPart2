package usecase

import "backend-b-payment-monitoring/models"

type RoleUsecaseInterface interface {
	GetAllRoleUsecase() ([]models.Role, error)
	AddRoleUsecase(role models.Role) (models.Role, error)
	GetRoleUsecase(id string) models.Role
	UpdateRoleUsecase(role models.Role, id string) (models.Role, error)
	DeleteRoleUsecase(role models.Role, id string) models.Role
}

type AccountUsecaseInterface interface {
	AddAccountUsecase(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error)
	GetAllAccountUsecase() ([]models.Role, error)
	GetAccountUsecase(account models.Role, id string) models.Role
	UpdateAccountUsecase(account models.Role, id string) (models.Role, error)
	DeleteAccountUsecase(account models.Role, id string) models.Role
}

type LoginUsecaseInterface interface {
	GetAdminId(name string) (models.Role, error)
	GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error)
	CreateAuth(user models.OfficerAccount) (*models.Auth, error)
	SignIn(authD models.Auth) (string, error)
	DeleteAuth(uuid string) error
}

type PaymentUsecaseInterface interface {
	GetAllRequestUsecase() ([]models.EntityPayment, error)
	GetReqByunit (id string) models.EntityPayment
	AddStatusReqUsecase(status models.StatusRequest)  models.StatusRequest
	UpdateAprovalUsecase(aproval models.EntityPayment, tipeReqRes string) (models.EntityPayment, error)
}

//type StatusReqUscInterface interface {
//	AddStatusReqUsecase(status models.StatusRequest) models.StatusRequest
//}
