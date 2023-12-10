package controller_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"petplace/back-mascotas/cmd/app/controller"
	"petplace/back-mascotas/cmd/app/controller/pet_errors"
	"petplace/back-mascotas/cmd/app/data"
	"petplace/back-mascotas/cmd/app/routes"
	"petplace/back-mascotas/cmd/app/services"
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
	require.NoError(t, err, fmt.Sprintf("body response: %s", w.Body))

	assert.Equal(t, expectedPet.Name, fetchedPet.Name)
	assert.Equal(t, expectedPet.Type.Normalice(), fetchedPet.Type)
	assert.Equal(t, expectedPet.BirthDate, fetchedPet.BirthDate)
	assert.Equal(t, expectedPet.OwnerID, fetchedPet.OwnerID)

	// In order to assert the creation date with no external clock, we'll verify the next:
	assert.LessOrEqual(t, expectedPet.RegisterDate, fetchedPet.RegisterDate)
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

//
// NewPet
//

func TestNewPetController_BadRequest(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	body := `{"Name":[\]}`
	response := pet_errors.EntityFormatError

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_EmptyRequest(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	body := `{}`
	response := pet_errors.EntityFormatError

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_InvalidAnimalType(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	badType := "Licha"
	body := fmt.Sprintf(`{"name": "Raaida", "type": "%s", "birth_date": "2013-05-25", "owner_id": "owner"}`, badType)
	response := pet_errors.InvalidAnimalType

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_InvalidBirthDate(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	badDate := "2000-13-01"
	body := fmt.Sprintf(`{"name": "Raaida", "type": "dog", "birth_date": "%s", "owner_id": "owner"}`, badDate)
	response := pet_errors.InvalidBirthDate

	w, req := newRequest(http.MethodPost, body, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestNewPetController_HappyPath(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	strDate := "2006-01-02"
	birthDate, err := time.Parse(time.DateOnly, strDate)
	require.NoError(t, err)

	requestedPet := data.Pet{
		Name:      "Pantufla",
		Type:      "cat",
		BirthDate: birthDate,
		OwnerID:   "tfanciotti",
	}
	var expectedPet data.Pet
	expectedPet.Name = requestedPet.Name
	expectedPet.Type = requestedPet.Type
	expectedPet.ID = requestedPet.ID
	expectedPet.BirthDate = requestedPet.BirthDate
	expectedPet.RegisterDate = time.Now()

	rawMsg := fmt.Sprintf(`{"name": "%s", "type": "%s", "birth_date": "%s", "owner_id": "%s"}`,
		requestedPet.Name,
		requestedPet.Type,
		strDate,
		requestedPet.OwnerID)

	petPlaceMock.EXPECT().RegisterNewPet(requestedPet).Return(expectedPet, nil)
	w, req := newRequest(http.MethodPost, rawMsg, "/pets/", nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertPet(t, http.StatusCreated, expectedPet, w)
}

//
// GetPet
//

func TestGetPetController_InvalidPetId(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	petID := "1234.0" // It must be an integer
	response := pet_errors.MissingParams

	url := fmt.Sprintf("/pets/%v", petID)
	w, req := newRequest(http.MethodGet, "", url, nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusBadRequest, response, w)
}

func TestGetPetController_NotFound(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	petID := 1234
	response := pet_errors.PetNotFound

	url := fmt.Sprintf("/pets/%d", petID)
	petPlaceMock.EXPECT().GetPet(petID).Return(data.Pet{}, nil)
	w, req := newRequest(http.MethodGet, "", url, nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusNotFound, response, w)
}

func TestGetPetController_HappyPath(t *testing.T) {
	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	strDate := "2006-01-02"
	birthDate, err := time.Parse(time.DateOnly, strDate)
	require.NoError(t, err)

	expectedPet := data.Pet{
		ID:           1234,
		Name:         "Pantufla",
		Type:         "cat",
		RegisterDate: time.Now(),
		BirthDate:    birthDate,
		OwnerID:      "tfanciotti",
	}

	url := fmt.Sprintf("/pets/%d", expectedPet.ID)
	petPlaceMock.EXPECT().GetPet(expectedPet.ID).Return(expectedPet, nil)
	w, req := newRequest(http.MethodGet, "", url, nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertPet(t, http.StatusOK, expectedPet, w)
}
