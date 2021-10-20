package repository

import (
	"backend-b-payment-monitoring/models"
)

type RoleRepoInterface interface {
	GetAllRoleRepo() ([]models.Role, error)
	AddRoleRepo(role models.Role) models.Role
	GetRoleRepo(id string) models.Role
	UpdateRoleRepo(role models.Role, id string) (models.Role, error)
	DeleteRoleRepo(role models.Role, id string) models.Role
	GetRoleRepoByName(name string) models.Role
}

type AccountRepoInterface interface {
	AddAccountRepo(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error)
	GetAllAccountRepo() ([]models.Role, error)
	GetAccountRepo(role models.Role, id string) models.Role
	UpdateAccountRepo(role models.Role, id string) (models.Role, error)
	DeleteAccountRepo(role models.Role, id string) models.Role
	GetAllAccountWorkUnit(username string) (int, error)
	GetAllAccountOfficer(username string) (int, error)
}

type LoginRepoInterface interface {
	GetAdminId(name string) (models.Role, error)
	GetDataWorkUnit(username string, password string) (models.WorkUnitAccount,error)
	CreateAuth(username string, roleId, roleName string) (*models.Auth, error)
	DeleteAuth (uuid string) error
	GetDataOfficerAccount(username string, password string, loginAs string) (models.OfficerAccount, error)
}

type PaymentRepoInterface interface {
	GetAllRequest() ([]models.EntityPayment, error)
	GetReqByunit(id string)([]models.EntityPayment, error)
	UpdateAprovalRepo(approval models.EntityPayment) (models.EntityPayment, error)
	AddStatusReqRepo(status models.StatusRequest) models.StatusRequest
	GetReqByIdPayment(id int) (models.EntityPayment, error)
}

//type StatusReqRepoInterface interface {
//	AddStatusReqRepo(status models.StatusRequest) models.StatusRequest
//}
