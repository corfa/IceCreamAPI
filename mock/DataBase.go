package mock

import (
	"iceCreamApiWithDI/database/ModelForGorm"
	"iceCreamApiWithDI/handler/ModelForGin"
)

type MockDataBase struct {
}

func (b MockDataBase) CreateUser(data ModelForGin.CreateUser) error {
	return nil
}
func (b MockDataBase) GetUser(user ModelForGin.GetUser) (*ModelForGorm.Users, error) {
	testUser := new(ModelForGorm.Users)
	return testUser, nil
}
func (b MockDataBase) GetIceCream() {

}
func (b MockDataBase) UpdateUser() {

}
func (b MockDataBase) DeleteIceCream() {

}
func (b MockDataBase) CreateIceCream() {

}
