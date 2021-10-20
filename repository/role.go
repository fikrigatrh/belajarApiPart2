package repository

import (
	"backend-b-payment-monitoring/models"

	"gorm.io/gorm"
)

type RoleRepoStruct struct {
	db *gorm.DB
}

func NewRoleRepoImpl(db *gorm.DB) RoleRepoInterface {
	return &RoleRepoStruct{db: db}
}

func (r *RoleRepoStruct) AddRoleRepo(role models.Role) models.Role {
	trx := r.db.Begin()
	err := trx.Debug().Create(&role).Error
	if err != nil {
		trx.Rollback()
		return models.Role{}
	}

	trx.Commit()

	return role
}

func (r RoleRepoStruct) GetAllRoleRepo() ([]models.Role, error) {
	var role []models.Role
	err := r.db.Debug().Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepoStruct) GetRoleRepo(id string) models.Role {
	var role models.Role
	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}

func (r *RoleRepoStruct) UpdateRoleRepo(role models.Role, id string) (models.Role, error) {
	err := r.db.Debug().Model(&role).Where("id = ?", id).Updates(&role).Error
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (r *RoleRepoStruct) DeleteRoleRepo(role models.Role, id string) models.Role {
	err := r.db.Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}

func (r *RoleRepoStruct) GetRoleRepoByName(name string) models.Role {
	var role models.Role
	err := r.db.Where("role_name = ?", name).First(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}