package models

import (
	"github.com/aaqyaar/telemedicine/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

func (u *User) BeforeSave(
	db *gorm.DB,
) error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
