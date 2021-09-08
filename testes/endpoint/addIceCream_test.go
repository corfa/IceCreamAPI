package endpoint

import (
	"iceCreamApiWithDI/layers/handler"
	"iceCreamApiWithDI/mock"
	"testing"
)


func TestHandlerAddIceCream_positive(t *testing.T) {
	Service := mock.MockService{}
	MyHandler := handler.NewHandler(Service)
	router := MyHandler.InitHandler()
	body := `{"Flavor":"testFlavor"}`
	w := mock.PerformRequest(router, "POST", "/ice-cream/add", body)
	if w.Code != 200 {
		t.Fatalf("Expected status code 200, got %v", w.Code)
	}

}
func TestHandlerAddIceCream_negative(t *testing.T) {
	Service := mock.MockService{}
	MyHandler := handler.NewHandler(Service)
	router := MyHandler.InitHandler()
	body := `{"wrong":"req"}`
	w := mock.PerformRequest(router, "POST", "/ice-cream/add", body)
	if w.Code != 400 {
		t.Fatalf("Expected status code 400, got %v", w.Code)
	}
}

