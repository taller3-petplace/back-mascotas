package controller_test

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"petplace/back-mascotas/cmd/app/routes"
	"strings"
	"testing"
)

func newRequest(method string, body string, url string, header http.Header) (*httptest.ResponseRecorder, *http.Request) {

	m := strings.ToUpper(method)
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(m, url, bodyReader)
	req.Header = header

	w := httptest.NewRecorder()

	return w, req
}

func assertHTTPResponse(t *testing.T, expeceteStatus int, expectedResponse string, w *httptest.ResponseRecorder) {

	assert.Equal(t, expeceteStatus, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())

}

func TestPing(t *testing.T) {

	mockRouter := routes.NewMockRouter()
	mockRouter.AddPingRoute()

	response := `{"message":"pong"}`

	w, req := newRequest(http.MethodGet, "", "/ping", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertHTTPResponse(t, http.StatusOK, response, w)

}
