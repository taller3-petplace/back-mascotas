package requester

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"petplace/back-mascotas/src/requester/domain"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func NewMockHttpClient() *MockHttpClient {
	return &MockHttpClient{}
}

func NewHttpClient() HttpClient {
	return &http.Client{}
}

func (m *MockHttpClient) Do(_ *http.Request) (*http.Response, error) {

	validBody := []domain.VaccineResponse{
		{
			ID:   "1",
			Name: "vacuna 1",
			Date: dateParser("2021-01-01"),
		},
		{
			ID:   "2",
			Name: "vacuna 2",
			Date: dateParser("2024-01-01"),
		},
		{
			ID:   "3",
			Name: "vacuna 3",
			Date: dateParser("2023-10-01"),
		},
		//{
		//	ID:   4,
		//	Name: "vacuna 4",
		//	Date: dateParser("2023-09-01"),
		//},
		{
			ID:   "2",
			Name: "vacuna 5",
			Date: dateParser("2023-08-01"),
		},
		{
			ID:   "6",
			Name: "vacuna 6",
			Date: dateParser("2023-07-01"),
		},
	}

	rawBody, _ := json.Marshal(validBody)
	bodyReader := bytes.NewReader(rawBody)

	readCloserBody := ioutil.NopCloser(bodyReader)
	return &http.Response{
		StatusCode: 200,
		Body:       readCloserBody,
		Header:     http.Header{},
	}, nil
}

func dateParser(date string) time.Time {
	parsedTime, _ := time.Parse(time.DateOnly, date)
	return parsedTime
}
