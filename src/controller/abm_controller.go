package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"petplace/back-mascotas/services"
	"strconv"
)

const IDParamName = "id"

type Entity interface {
	IsZeroValue() bool
}

type ABMController[T Entity] struct {
	s        services.ABMService[T]
	Validate func(T) error
}

func NewABMController[T Entity](service services.ABMService[T]) ABMController[T] {
	return ABMController[T]{s: service}
}

func (abm *ABMController[Entity]) New(c *gin.Context) {

	var e Entity
	err := c.BindJSON(&e)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	err = abm.Validate(e)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, ValidationError, err.Error())
		return
	}

	e, err = abm.s.New(e)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, e)
}

func (abm *ABMController[Entity]) Get(c *gin.Context) {

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := abm.s.Get(id)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if e.IsZeroValue() {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}

	c.JSON(http.StatusOK, e)
}

func (abm *ABMController[Entity]) Edit(c *gin.Context) {
	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := abm.s.Get(id)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}
	if e.IsZeroValue() {
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}

	err = c.BindJSON(&e)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	err = abm.Validate(e)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, ValidationError, err.Error())
		return
	}

	e, err = abm.s.Edit(id, e)
	if err != nil {
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, e)

}

func (abm *ABMController[Entity]) Delete(c *gin.Context) {

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	abm.s.Delete(id)

	c.JSON(http.StatusNoContent, nil)
}
