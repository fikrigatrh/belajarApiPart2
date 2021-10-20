package repository

import (
	"backend-b-payment-monitoring/models"

	"gorm.io/gorm"
)

type AccountRepoStruct struct {
	db *gorm.DB
}

func NewAccountRepoImpl(db *gorm.DB) AccountRepoInterface {
	return &AccountRepoStruct{db: db}
}

func (r *AccountRepoStruct) AddAccountRepo(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error) {
	accCustomer := models.WorkUnitAccount{
		Name:     accOfficer.Name,
		Username: accOfficer.Username,
		Password: accOfficer.Password,
	}

	if accOfficer.Role == 4 {
		err := r.db.Create(&accCustomer).Error
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return models.OfficerAccount{}, accCustomer, err
	} else {
		err := r.db.Create(&accOfficer).Error
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return accOfficer, models.WorkUnitAccount{}, err
	}

}

func (r AccountRepoStruct) GetAllAccountRepo() ([]models.Role, error) {
	var account []models.Role
	err := r.db.Debug().Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepoStruct) GetAccountRepo(account models.Role, id string) models.Role {
	err := r.db.Where("id = ?", id).First(&account).Error
	if err != nil {
		return models.Role{}
	}

	return account
}

func (r *AccountRepoStruct) UpdateAccountRepo(account models.Role, id string) (models.Role, error) {
	err := r.db.Debug().Model(&account).Where("id = ?", id).Updates(&account).Error
	if err != nil {
		return models.Role{}, err
	}

	return account, nil
}

func (r *AccountRepoStruct) DeleteAccountRepo(account models.Role, id string) models.Role {
	err := r.db.Where("id = ?", id).Delete(&account).Error
	if err != nil {
		return models.Role{}
	}

	return account
}

func (r AccountRepoStruct) GetAllAccountWorkUnit(username string) (int, error) {
	var count int64
	var tbl models.WorkUnitAccount
	err := r.db.Debug().Model(tbl).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}


func (r AccountRepoStruct) GetAllAccountOfficer(username string) (int, error) {
	var count int64
	var tbl models.OfficerAccount
	err := r.db.Debug().Model(tbl).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}