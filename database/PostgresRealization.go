package database

import (
	"errors"
	"gorm.io/gorm"
	"iceCreamApiWithDI/database/ModelForGorm"
	"iceCreamApiWithDI/handler/ModelForGin"
)

type PostgresDB struct {
	DB *gorm.DB
}

func (u *PostgresDB) CreateUser(data ModelForGin.CreateUser) error {

	DBresult := u.DB.Create(&ModelForGorm.Users{Login: data.Login, Password: data.Password, Email: data.Email})
	return DBresult.Error

}

func (u *PostgresDB) GetUser(data ModelForGin.GetUser) (*ModelForGorm.Users, error) {
	user := new(ModelForGorm.Users)
	errDB := u.DB.Where("login = ?", data.Login).Find(&user)
	if errDB.RowsAffected == 0 {
		return user, errors.New("User not found")
	}
	return user, nil
}
func (u *PostgresDB) UpdateUser() {

}
func (u *PostgresDB) CreateIceCream() {

}

func (u *PostgresDB) DeleteIceCream() {

}
func (u *PostgresDB) GetIceCream() {

}

func NewPostgres(db *gorm.DB) *PostgresDB {
	return &PostgresDB{
		DB: db,
	}
}
