package controller

import (
	"github.com/gin-gonic/gin"
)

type APIError struct {
	error
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewApiError(err error, status int) APIError {
	return APIError{
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
