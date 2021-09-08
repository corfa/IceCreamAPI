package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Sslmode  string
}

func EnginPostgres(config ConfigDB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.Port, config.Sslmode)), &gorm.Config{})
	if err != nil {
		return nil, errors.New("error connection")
	}
	return db, err

}
