package controller_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/controller/pet_errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/routes"
	"strings"
	"testing"
	"time"
)

func newRequest(method string, body string, url string, header http.Header) (*httptest.ResponseRecorder, *http.Request) {

	m := strings.ToUpper(method)
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(m, url, bodyReader)
	req.Header = header

	w := httptest.NewRecorder()

	return w, req
}

func assertHTTPResponse(t *testing.T, expectedStatus int, expectedResponse string, w *httptest.ResponseRecorder) {

	assert.Equal(t, expectedStatus, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())

}

func assertPet(t *testing.T, status int, expectedPet data.Pet, w *httptest.ResponseRecorder) {

	assert.Equal(t, status, w.Code)

	var fetchedPet data.Pet
	err := json.Unmarshal(w.Body.Bytes(), &fetchedPet)
	require.NoError(t, err)

	assert.Equal(t, expectedPet.Name, fetchedPet.Name)
	assert.Equal(t, expectedPet.Type.Normalice(), fetchedPet.Type)
	assert.Equal(t, expectedPet.BirthDate, fetchedPet.BirthDate)
	assert.Equal(t, expectedPet.Owner, fetchedPet.Owner)

	// In order to assert the creation date with no external clock, we'll verify the next:
	assert.Less(t, expectedPet.RegisterDate, fetchedPet.RegisterDate)
	assert.Less(t, fetchedPet.RegisterDate, time.Now())
}

func assertError(t *testing.T, expectedStatus int, expectedError error, w *httptest.ResponseRecorder) {

	fetched := controller.APIError{}

	err := json.Unmarshal(w.Body.Bytes(), &fetched)
	require.NoError(t, err)

	assert.Equal(t, expectedStatus, w.Code)
	assert.Contains(t, fetched.Message, expectedError.Error(), fmt.Sprintf("BODY: %s  ", w.Body))
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

func TestNewPetController_BadRequest(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	err := mockRouter.AddPetRoutes()
	require.NoError(t, err)

	body := `{"Name":[\]}`
	response := pet_errors.EntityFormatError

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_InvalidAnimalType(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	err := mockRouter.AddPetRoutes()
	require.NoError(t, err)

	body := `{"type":"Licha"}`
	response := pet_errors.InvalidAnimalType

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_HappyPath(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	err := mockRouter.AddPetRoutes()
	require.NoError(t, err)

	strDate := "2006-01-02T15:04:05-04:00"
	birthDate, err := time.Parse(time.RFC3339, strDate)
	require.NoError(t, err)

	expectedPet := data.Pet{
		Name:         "Pantufla",
		Type:         "cat",
		RegisterDate: time.Now(),
		BirthDate:    birthDate,
		Owner:        "tfanciotti",
	}

	rawMsg := fmt.Sprintf(`{"name": "%s", "type": "%s", "birth_date": "%s", "owner": "%s"}`,
		expectedPet.Name,
		expectedPet.Type,
		strDate,
		expectedPet.Owner)

	w, req := newRequest(http.MethodPost, rawMsg, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertPet(t, http.StatusCreated, expectedPet, w)
}
