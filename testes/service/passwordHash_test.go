package service

import (
	"iceCreamApiWithDI/mock"
	"iceCreamApiWithDI/service"
	"testing"
)

func TestPasswordHash_positive(t *testing.T) {
	DataBase := mock.MockDataBase{}
	Service := service.NewServices(DataBase)
	testPassword := "qwerty"
	hashPass := Service.HashPassword(testPassword)
	if Service.CheckHashPassword(testPassword, hashPass) != true {
		t.Fatalf("Expected true, got %v", Service.CheckHashPassword(testPassword, hashPass))
	}

}
func TestPasswordHash_negative(t *testing.T) {
	DataBase := mock.MockDataBase{}
	Service := service.NewServices(DataBase)
	testPassword := "qwerty"
	hashPass := Service.HashPassword(testPassword)
	testPassword = "wrongQwerty"
	if Service.CheckHashPassword(testPassword, hashPass) != false {
		t.Fatalf("Expected false, got %v", Service.CheckHashPassword(testPassword, hashPass))
	}
}
