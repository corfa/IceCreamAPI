package service

import (
	"iceCreamApiWithDI/layers/service"
	"iceCreamApiWithDI/mock"
	"testing"
)

func TestJWToken_positive(t *testing.T) {
	DataBase := mock.MockDataBase{}
	Service := service.NewServices(DataBase)
	TestID := 1
	token, errToken := Service.CreateJWToken(TestID)
	if errToken != nil {
		t.Fatalf("error in token not null,  %v", errToken)
	}
	OutID, errReadToken := Service.ReadJWToken(token)
	if errReadToken != nil {
		t.Fatalf("error in token read not null,  %v", errReadToken)
	}
	if OutID != TestID {
		t.Fatalf("Expected 1, got %v", OutID)
	}

}
