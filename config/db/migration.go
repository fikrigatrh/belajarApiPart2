package db

import (
	"backend-b-payment-monitoring/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	var err error

	//db.Debug().Migrator().DropTable(models.EntityPayment{})
	// isi param automigrate dengan model/entitas
	err = db.Debug().AutoMigrate(
		models.OfficerAccount{},
		models.WorkUnitAccount{},
		models.Role{},
		models.Auth{},
		models.AuthDetail{},
		models.EntityPayment{},
	)
	if err != nil {
		return
	}
}
