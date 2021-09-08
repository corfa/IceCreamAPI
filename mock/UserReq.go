package mock

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func PerformRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	testBody := ioutil.NopCloser(strings.NewReader(body))
	req.Body = testBody
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
