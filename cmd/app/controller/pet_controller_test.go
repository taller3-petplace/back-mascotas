package controller_test

import (
	"encoding/json"
	"errors"
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

func assertSearchResult(t *testing.T, status int, expectedResponse data.SearchResponse, w *httptest.ResponseRecorder) {

	assert.Equal(t, status, w.Code)

	var fetched data.SearchResponse
	err := json.Unmarshal(w.Body.Bytes(), &fetched)
	require.NoError(t, err, fmt.Sprintf("body response: %s", w.Body))

	assert.Equal(t, expectedResponse.Paging.Offset, fetched.Paging.Offset)
	assert.Equal(t, expectedResponse.Paging.Limit, fetched.Paging.Limit)
	assert.Equal(t, expectedResponse.Paging.Total, fetched.Paging.Total)
	assert.LessOrEqual(t, len(expectedResponse.Results), len(fetched.Results))
	for _, petExpected := range expectedResponse.Results {
		exist := false
		for _, petFetched := range fetched.Results {
			exist = exist || (petExpected.OwnerID == petFetched.OwnerID &&
				petExpected.Name == petFetched.Name &&
				petExpected.Type == petFetched.Type)
		}
		assert.True(t, exist)
	}
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

	w, req := newRequest(http.MethodPost, body, "/pets/pet", nil)
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

	w, req := newRequest(http.MethodPost, body, "/pets/pet", nil)
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

	w, req := newRequest(http.MethodPost, body, "/pets/pet", nil)
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

	w, req := newRequest(http.MethodPost, body, "/pets/pet", nil)
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
	w, req := newRequest(http.MethodPost, rawMsg, "/pets/pet", nil)
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

	url := fmt.Sprintf("/pets/pet/%v", petID)
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

	url := fmt.Sprintf("/pets/pet/%d", petID)
	petPlaceMock.EXPECT().GetPet(petID).Return(data.Pet{}, nil)
	w, req := newRequest(http.MethodGet, "", url, nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertError(t, http.StatusNotFound, response, w)
}

func TestGetPetController_ServiceError(t *testing.T) {

	petID := 1234
	url := fmt.Sprintf("/pets/pet/%d", petID)

	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	expectedError := errors.New("this is a simulated error")

	w, req := newRequest(http.MethodGet, "", url, nil)
	petPlaceMock.EXPECT().GetPet(petID).Return(data.Pet{}, expectedError)
	mockRouter.ServeRequest(w, req)

	assertError(t, http.StatusInternalServerError, expectedError, w)
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

	url := fmt.Sprintf("/pets/pet/%d", expectedPet.ID)
	petPlaceMock.EXPECT().GetPet(expectedPet.ID).Return(expectedPet, nil)
	w, req := newRequest(http.MethodGet, "", url, nil)
	mockRouter.ServeRequest(w, req)

	//Assertion
	assertPet(t, http.StatusOK, expectedPet, w)
}

//
// GetPetsByOwner
//

func TestGetPetsByOwnerController_InvalidQueryParams(t *testing.T) {

	testCases := []string{
		//"/pets/owner/tfanciotti?offset=&limit",
		"/pets/owner/tfanciotti?offset=X&limit=Y",
		"/pets/owner/tfanciotti?offset=10&limit=Y",
		"/pets/owner/tfanciotti?offset=X&limit=20",
		//"/pets/owner/tfanciotti?offset=10&limit=20",
	}

	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	for i, url := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			w, req := newRequest(http.MethodGet, "", url, nil)
			mockRouter.ServeRequest(w, req)
			assertError(t, http.StatusBadRequest, pet_errors.MissingParams, w)
		})
	}
}

func TestGetPetsByOwnerController_NotFound(t *testing.T) {

	ownerID := "tfanciotti"
	url := fmt.Sprintf("/pets/owner/%s", ownerID)

	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	searchReq := data.NewSearchRequest()
	searchReq.OwnerId = ownerID

	w, req := newRequest(http.MethodGet, "", url, nil)
	petPlaceMock.EXPECT().GetPetsByOwner(searchReq).Return(data.SearchResponse{
		Paging:  data.Paging{},
		Results: nil,
	}, nil)
	mockRouter.ServeRequest(w, req)
	assertError(t, http.StatusNotFound, pet_errors.PetNotFound, w)

}

func TestGetPetsByOwnerController_ServiceError(t *testing.T) {

	ownerID := "tfanciotti"
	url := fmt.Sprintf("/pets/owner/%s", ownerID)

	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	searchReq := data.NewSearchRequest()
	searchReq.OwnerId = ownerID

	expectedError := errors.New("this is a simulated error")

	w, req := newRequest(http.MethodGet, "", url, nil)
	petPlaceMock.EXPECT().GetPetsByOwner(searchReq).Return(data.SearchResponse{}, expectedError)
	mockRouter.ServeRequest(w, req)
	assertError(t, http.StatusInternalServerError, expectedError, w)

}

func TestGetPetsByOwnerController_HappyPath(t *testing.T) {

	ownerID := "tfanciotti"
	baseUrl := fmt.Sprintf("/pets/owner/%s", ownerID)

	var allPetsOfTomi = []data.Pet{
		{
			ID:           1234,
			Name:         "Raaida",
			Type:         "dog",
			RegisterDate: time.Now(),
			BirthDate:    time.Time{},
			OwnerID:      ownerID,
		},
		{
			ID:           12345,
			Name:         "Javo",
			Type:         "cat",
			RegisterDate: time.Now(),
			BirthDate:    time.Time{},
			OwnerID:      ownerID,
		},
	}

	var testCases = []struct {
		Name   string
		Result []data.Pet
		Url    string
		Owner  string
		Total  uint
		Offset uint
		Limit  uint
	}{
		{
			Name:   "Both pets of Tomi (without query params)",
			Result: allPetsOfTomi,
			Url:    baseUrl,
			Total:  2,
			Offset: 0,
			Limit:  10,
			Owner:  "tfanciotti",
		},
		{
			Name:   "First pet (limit=1)",
			Url:    baseUrl + "?limit=1",
			Result: []data.Pet{allPetsOfTomi[0]},
			Total:  2,
			Limit:  1,
			Offset: 0,
			Owner:  "tfanciotti",
		},
		{
			Name:   "Second pet (offset=1)",
			Url:    baseUrl + "?offset=1",
			Result: []data.Pet{allPetsOfTomi[1]},
			Total:  2,
			Limit:  10,
			Offset: 1,
			Owner:  "tfanciotti",
		},
	}

	mockRouter := routes.NewMockRouter()
	petPlaceMock := services.NewMockPetService(gomock.NewController(t))
	err := mockRouter.AddPetRoutes(petPlaceMock)
	require.NoError(t, err)

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {

			searchReq := data.NewSearchRequest()
			searchReq.OwnerId = ownerID
			searchReq.Limit = testCase.Limit
			searchReq.Offset = testCase.Offset

			var expectedResult = data.SearchResponse{
				Paging: data.Paging{
					Total:  testCase.Total,
					Offset: testCase.Offset,
					Limit:  testCase.Limit,
				},
				Results: testCase.Result,
			}

			w, req := newRequest(http.MethodGet, "", testCase.Url, nil)
			petPlaceMock.EXPECT().GetPetsByOwner(searchReq).Return(expectedResult, nil)
			mockRouter.ServeRequest(w, req)

			assertSearchResult(t, http.StatusOK, expectedResult, w)
		})
	}
}
