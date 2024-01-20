package models

type User struct {
	id       string `gorm:"primaryKey"`
	name     string
	email    string `gorm:"uniqueIndex"`
	password string
}
