package mock

import (
	"iceCreamApiWithDI/layers/database/ModelForGorm"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
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


func (b MockDataBase) DeleteIceCream(iceCream ModelForGin.IceCream) error {
	return nil
}
func (b MockDataBase) CreateIceCream(iceCream ModelForGin.IceCream) error {
	return nil

}
