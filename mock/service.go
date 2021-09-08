package mock

import "iceCreamApiWithDI/handler/ModelForGin"

type MockService struct {
}

func (s MockService) CreateUser(user ModelForGin.CreateUser) error {
	return nil

}
func (s MockService) HashPassword(str string) string {
	return "testBody"
}
func (s MockService) CheckHashPassword(pass string, hash string) bool {
	return true
}
func (s MockService) CreateJWToken() {

}
func (s MockService) ReadJWToken() {

}
