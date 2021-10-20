package paymentRepo

import (
	"backend-b-payment-monitoring/models"
	"backend-b-payment-monitoring/repository"
	"gorm.io/gorm"
)

type PaymentRepoStruct struct {
	db *gorm.DB
}

func NewPaymentImpl(db *gorm.DB) repository.PaymentRepoInterface {
	return &PaymentRepoStruct{db}
}

func (pr PaymentRepoStruct) GetAllRequest() ([]models.EntityPayment, error) {
	var request []models.EntityPayment
	err := pr.db.Debug().Find(&request).Error
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (pr PaymentRepoStruct ) GetReqByunit(id string) models.EntityPayment  {
	var transaksi models.EntityPayment
	err := pr.db.Where("id_unit = ?", id).First(&transaksi).Error
	if err != nil {
		return models.EntityPayment{}
	}
	return transaksi
}

//func (ar PaymentRepoStruct) UpdateAprovalRepo(approval models.Approval, id string) (models.Approval, error)  {
//	//err := ar.db.Debug().Model(&approval).Where("id", id).Updates(&approval).Error
//	//if err != nil {
//	//	return models.Approval{},err
//	//}
//	//return approval, nil
//}

func (pr *PaymentRepoStruct) AddStatusReqRepo(status models.StatusRequest) models.StatusRequest{
	err := pr.db.Create(&status).Error
	if err != nil {
		return models.StatusRequest{}
	}
	return status
}

func (pr PaymentRepoStruct ) GetReqByIdPayment(id int) (models.EntityPayment, error)  {
	var trans models.EntityPayment
	err := pr.db.Debug().Where("id_payment = ?", id).First(&trans).Error
	if err != nil {
		return models.EntityPayment{}, err
	}
	return trans, nil
}

func (pr PaymentRepoStruct) UpdateAprovalRepo(approval models.EntityPayment) (models.EntityPayment, error)  {
	trx := pr.db.Begin()

	err := pr.db.Debug().Save(&approval).Error
	if err != nil {
		trx.Rollback()
		return models.EntityPayment{}, err
	}

	trx.Commit()
	return approval, nil
}