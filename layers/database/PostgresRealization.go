package database

import (
	"errors"
	"gorm.io/gorm"
	"iceCreamApiWithDI/layers/database/ModelForGorm"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
)

type PostgresDB struct {
	DB *gorm.DB
}

func (u *PostgresDB) CreateUser(data ModelForGin.CreateUser) error {

	Queries := u.DB.Create(&ModelForGorm.Users{Login: data.Login, Password: data.Password, Email: data.Email})
	if Queries.Error!=nil{
		return Queries.Error
	}
	return nil

}

func (u *PostgresDB) GetUser(data ModelForGin.GetUser) (*ModelForGorm.Users, error) {
	user := new(ModelForGorm.Users)
	Queries := u.DB.Where("login = ?", data.Login).Find(&user)
	if Queries.RowsAffected == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}

func (u *PostgresDB) CreateIceCream(iceCream ModelForGin.IceCream)error {
	Queries:=u.DB.Create(&ModelForGorm.IceCreams{Flavor: iceCream.Flavor})
	if Queries.Error!=nil{
		return Queries.Error
	}
	return nil
}

func (u *PostgresDB) DeleteIceCream(IceCream ModelForGin.IceCream) error {
	Queries:=u.DB.Where("flavor = ?",IceCream.Flavor).Delete(&ModelForGorm.IceCreams{})
	if Queries.Error!=nil{
		return Queries.Error
	}
	return nil
}


func NewPostgres(db *gorm.DB) *PostgresDB {
	return &PostgresDB{
		DB: db,
	}
}
