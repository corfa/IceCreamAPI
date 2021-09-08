package mock

import (
	ModelForGin2 "iceCreamApiWithDI/layers/handler/ModelForGin"
)

type MockService struct {
}

func (s MockService) CreateUser(user ModelForGin2.CreateUser) error {
	return nil

}
func (s MockService) LoginUser(user ModelForGin2.GetUser) (string, error)  {
	return "",nil
}
func (s MockService) HashPassword(str string) string {
	return "testBody"
}
func (s MockService) AppendIceCream(IceCream ModelForGin2.IceCream,token ModelForGin2.HeaderAuth)error  {
	return nil
}
func (s MockService) DeleteIceCream(IceCream ModelForGin2.IceCream,token ModelForGin2.HeaderAuth)error  {
	return nil
}
func (s MockService) CheckHashPassword(pass string, hash string) bool {
	return true
}
func (s MockService) CreateJWToken(int) (string,error) {
	return "",nil
}

func (s MockService) ReadJWToken(string) (int,error) {
	return 0,nil
}
