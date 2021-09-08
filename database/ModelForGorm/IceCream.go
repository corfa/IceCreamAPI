package ModelForGorm

import "gorm.io/gorm"

type IceCreams struct {
	gorm.Model
	Id     int    `gorm:"primaryKey"`
	Flavor string `gorm:"unique"`
}
