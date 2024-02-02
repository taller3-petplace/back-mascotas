package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/src/model"
	"strconv"
)

const (
	offsetQueryParam = "offset"
	limitQueryParam  = "limit"

	offsetDefault = "0"
	limitDefault  = "10"
)

func getSearchParams(c *gin.Context) (*model.SearchParams, *APIError) {

	offset, err := strconv.Atoi(c.DefaultQuery(offsetQueryParam, offsetDefault))
	if err != nil {
		apiErr := fmt.Errorf("%v: %v", MissingParams, err.Error())
		return nil, NewApiError(apiErr, http.StatusBadRequest)
	}

	limit, err := strconv.Atoi(c.DefaultQuery(limitQueryParam, limitDefault))
	if err != nil {
		apiErr := fmt.Errorf("%v: %v", MissingParams, err.Error())
		return nil, NewApiError(apiErr, http.StatusBadRequest)
	}

	return &model.SearchParams{
		Offset: uint(offset),
		Limit:  uint(limit),
	}, nil
}

func getLocation(c *gin.Context) (*model.Location, *APIError) {
	latitudeStr := c.Query("latitude")
	if latitudeStr == "" {
		err := fmt.Errorf("%v: expected longitude", MissingParams)
		return nil, NewApiError(err, http.StatusBadRequest)
	}

	longitudeStr := c.Query("longitude")
	if longitudeStr == "" {
		err := fmt.Errorf("%v: expected longitude", MissingParams)
		return nil, NewApiError(err, http.StatusBadRequest)
	}

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		apiErr := fmt.Errorf("error parsing latitude: %v", err)
		return nil, NewApiError(apiErr, http.StatusBadRequest)
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		apiErr := fmt.Errorf("error parsing longitude: %v", err)
		return nil, NewApiError(apiErr, http.StatusBadRequest)
	}

	loc := model.Location{
		Latitude:  latitude,
		Longitude: longitude,
	}

	return &loc, nil
}
