package usecase

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/repository"
	"errors"
	gtr_validate "github.com/fikrigatrh/validate"
)

type AccountUsecaseStruct struct {
	account repository.AccountRepoInterface
}

func NewAccountUsecaseImpl(account repository.AccountRepoInterface) AccountUsecaseInterface {
	return &AccountUsecaseStruct{account: account}
}

func (u AccountUsecaseStruct) AddAccountUsecase(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error) {
	check := gtr_validate.Num(accOfficer.Username)
	if check {
		return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("username must be number")
	}

	if accOfficer.Role != models.Admin {
		if accOfficer.Role != models.GeneralSupport {
			if accOfficer.Role != models.Accounting {
				if accOfficer.Role != models.Customer {
					return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("role not available")
				}
			}
		}
	}

	if accOfficer.Role == 4 {
		count, err := u.account.GetAllAccountWorkUnit(accOfficer.Username)
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		} else if count > 0 {
			return models.OfficerAccount{}, models.WorkUnitAccount{},  errors.New("username has been registered")
		}
	} else {
		count, err := u.account.GetAllAccountOfficer(accOfficer.Username)
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		} else if count > 0 {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("username has been registered")
		}
	}

	accDataOfficer, accDataCustomer, err := u.account.AddAccountRepo(accOfficer)
	if err != nil {
		return models.OfficerAccount{}, models.WorkUnitAccount{}, err
	}

	return accDataOfficer, accDataCustomer, nil
}

func (u AccountUsecaseStruct) GetAccountUsecase(account models.Role, id string) models.Role {
	accountData := u.account.GetAccountRepo(account, id)

	return accountData
}

func (u AccountUsecaseStruct) UpdateAccountUsecase(account models.Role, id string) (models.Role, error) {
	accountData := u.account.GetAccountRepo(account, id)
	if accountData.ID == 0 {
		return models.Role{}, errors.New("error")
	}

	accountData, err := u.account.UpdateAccountRepo(account, id)
	if err != nil {
		return models.Role{}, err
	}
	return accountData, nil
}

func (u AccountUsecaseStruct) DeleteAccountUsecase(account models.Role, id string) models.Role {
	accountData := u.account.DeleteAccountRepo(account, id)

	return accountData
}

func (u AccountUsecaseStruct) GetAllAccountUsecase() ([]models.Role, error) {
	getAll, err := u.account.GetAllAccountRepo()
	if err != nil {
		return nil, err
	}

	return getAll, nil
}
