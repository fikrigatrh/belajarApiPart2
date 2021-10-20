package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName string `json:"role_name"`
}
