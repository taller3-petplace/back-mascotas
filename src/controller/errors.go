package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	EntityFormatError = errors.New("entity could not be mapped")
	ValidationError   = errors.New("invalid entity")
	RegisterError     = errors.New("error trying to register pet")
	InvalidAnimalType = errors.New("invalid animal type")
	InvalidBirthDate  = errors.New("invalid birth_date")
	MissingParams     = errors.New("missing params on request")
	EntityNotFound    = errors.New("entity not found")
	ServiceError      = errors.New("service error")
)

type APIError struct {
	error   `swaggerignore:"true"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewApiError(err error, status int) *APIError {
	return &APIError{
		error:   err,
		Status:  status,
		Message: err.Error(),
	}
}

func ReturnError(c *gin.Context, status int, err error, msg string) {
	c.JSON(status, APIError{
		error:   err,
		Status:  status,
		Message: err.Error() + ":" + msg,
	})
}
