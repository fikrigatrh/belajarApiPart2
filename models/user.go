package models

import (
	"time"
)

type OfficerAccount struct {
	ID        uint       `gorm:"primary_key" json:"id_user"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name     string `json:"nama"`
	LoginAs  string `json:"login_as,omitempty" gorm:"-"`
	Role     int    `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type WorkUnitAccount struct {
	ID        uint       `gorm:"primary_key" json:"id_user"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Name     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
}
