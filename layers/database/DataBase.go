package database

import (
	"gorm.io/gorm"
	"iceCreamApiWithDI/layers/database/ModelForGorm"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
)

type IDataBase interface {
	CreateUser(ModelForGin.CreateUser) error
	GetUser(user ModelForGin.GetUser) (*ModelForGorm.Users, error)
	CreateIceCream(iceCream ModelForGin.IceCream) error
	DeleteIceCream(IceCream ModelForGin.IceCream) error
}

type DataBase struct {
	IDataBase
}

func NewDataBase(DB *gorm.DB) *DataBase {
	return &DataBase{
		NewPostgres(DB),
	}
}
