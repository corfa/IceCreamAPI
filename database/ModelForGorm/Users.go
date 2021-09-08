package ModelForGorm

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id       int    `gorm:"primaryKey"`
	Login    string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
}
