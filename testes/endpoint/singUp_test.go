package endpoint

import (
	"iceCreamApiWithDI/layers/handler"
	"iceCreamApiWithDI/mock"
	"testing"
)

func TestHandlerSingUp_positive(t *testing.T) {
	Service := mock.MockService{}
	MyHandler := handler.NewHandler(Service)
	router := MyHandler.InitHandler()
	body := `{"password":"qwerty","login":"test", "email":"testemail"}`
	w := mock.PerformRequest(router, "POST", "/auth/sing-up", body)
	if w.Code != 200 {
		t.Fatalf("Expected status code 200, got %v", w.Code)
	}

}
func TestHandlerSingUp_negative(t *testing.T) {
	Service := mock.MockService{}
	MyHandler := handler.NewHandler(Service)
	router := MyHandler.InitHandler()
	body := `{"wrong":"req"}`
	w := mock.PerformRequest(router, "POST", "/auth/sing-up", body)
	if w.Code != 400 {
		t.Fatalf("Expected status code 400, got %v", w.Code)
	}
}
