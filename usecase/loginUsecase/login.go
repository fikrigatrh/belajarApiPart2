package loginUsecase

import (
	"backend-b-payment-monitoring/auth"
	"backend-b-payment-monitoring/config/env"
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/repository"
	"backend-b-payment-monitoring/usecase"
	"errors"
	"strconv"
)

type LoginUsecaseStruct struct {
	LoginRepo repository.LoginRepoInterface
}

func NewLoginUsecaseImpl(LoginRepo repository.LoginRepoInterface) usecase.LoginUsecaseInterface {
	return &LoginUsecaseStruct{LoginRepo}
}

func (lu LoginUsecaseStruct) GetAdminId(name string) (models.Role, error) {

	res, err := lu.LoginRepo.GetAdminId(name)
	if err != nil {
		return models.Role{}, err
	}
	return res, nil
}

func (lu LoginUsecaseStruct) GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error)  {
	res, err := lu.LoginRepo.GetDataWorkUnit(username, password)
	if err != nil {
		return models.WorkUnitAccount{}, err
	}
	return res, nil
}

func (lu LoginUsecaseStruct) CreateAuth(user models.OfficerAccount) (*models.Auth, error) {
	var idCust int
	var idOffice int

	roleTemp, err := strconv.Atoi(user.LoginAs)

	if roleTemp != 0 {
		if roleTemp != models.Admin {
			if roleTemp != models.GeneralSupport {
				if roleTemp != models.Accounting {
					if roleTemp != models.Customer {
						return nil, errors.New("role not available")
					}
				}
			}
		}
	}

	// melakukan pengecekan terhadap value login_as yg dikirim dari depan
	if user.LoginAs == "" { // login_as tidak ada isinya
		if user.Username == env.Config.UsernameSu && user.Password == env.Config.PasswordSu { // cek data dengan username & pass di env untuk super admin
			name := "admin"
			adminId, err := lu.LoginRepo.GetAdminId(name)
			if err != nil {
				return nil, err
			}
			user.LoginAs = adminId.RoleName
		} else {
			//check ke table work unit apakah data user sudah terdaftar atau belum dengan param username dan password
			customer, err := lu.LoginRepo.GetDataWorkUnit(user.Username, user.Password)
			if err != nil {
				return nil, err
			}
			// jika ada datanya, kita ambil id customernya, lalu ditampung ke variable idCust
			idCust = int(customer.ID)
			user.LoginAs = "work_unit"
		}
	} else {
		//check data ke table officer
		account, err := lu.LoginRepo.GetDataOfficerAccount(user.Username, user.Password, user.LoginAs)
		if err != nil {
			return nil, err
		}
		idOffice = int(account.ID)
		user.LoginAs = "officer_account"
	}

	dataAuth, err := lu.LoginRepo.CreateAuth(user.Username, user.Password, user.LoginAs)
	if err != nil {
		return nil, err
	}

	dataAuth.IdCust = idCust
	dataAuth.IdOffice = idOffice
	return dataAuth, nil
}

func (lu LoginUsecaseStruct) SignIn(authD models.Auth) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (lu LoginUsecaseStruct) DeleteAuth(uuid string) error {
	return lu.LoginRepo.DeleteAuth(uuid)
}