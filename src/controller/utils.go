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
