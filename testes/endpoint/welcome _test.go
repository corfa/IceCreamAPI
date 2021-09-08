package endpoint

import (
	"fmt"
	"iceCreamApiWithDI/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcome(t *testing.T) {
	h := new(handler.Handler)

	ts := httptest.NewServer(h.InitHandler())

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//newStr := buf.String()

}
