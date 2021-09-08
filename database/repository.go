package database

import (
	"gorm.io/gorm"
	"iceCreamApiWithDI/database/ModelForGorm"
	"iceCreamApiWithDI/handler/ModelForGin"
)

type IDataBase interface {
	CreateUser(ModelForGin.CreateUser) error
	GetUser(user ModelForGin.GetUser) (*ModelForGorm.Users, error)
	UpdateUser()
	CreateIceCream()
	DeleteIceCream()
	GetIceCream()
}

type DataBase struct {
	IDataBase
}

func NewDataBase(DB *gorm.DB) *DataBase {
	return &DataBase{
		NewPostgres(DB),
	}
}
