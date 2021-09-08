package service

import (
	"iceCreamApiWithDI/database"
	"iceCreamApiWithDI/handler/ModelForGin"
	"iceCreamApiWithDI/service/realizationService"
)

type IServices interface {
	CreateUser(user ModelForGin.CreateUser) error
	LoginUser(user ModelForGin.GetUser) (string, error)
	HashPassword(string) string
	CheckHashPassword(pass string, hash string) bool
	CreateJWToken(id int) (string, error)
	ReadJWToken(Token string) (int, error)
}
type Services struct {
	IServices
}

func NewServices(db database.IDataBase) *Services {
	return &Services{
		realizationService.NewService(db),
	}
}
