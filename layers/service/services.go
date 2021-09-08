package service

import (
	"iceCreamApiWithDI/layers/database"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
	"iceCreamApiWithDI/layers/service/realizationService"
)

type IServices interface {
	CreateUser(user ModelForGin.CreateUser) error
	LoginUser(user ModelForGin.GetUser) (string, error)

	AppendIceCream(iceCream ModelForGin.IceCream, token ModelForGin.HeaderAuth) error
	DeleteIceCream(iceCream ModelForGin.IceCream, token ModelForGin.HeaderAuth) error

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
