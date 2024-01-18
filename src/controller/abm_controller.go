package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"petplace/back-mascotas/src/services"
	"strconv"
	"strings"
)

const IDParamName = "id"

const logTemplate = "ABMController: %s | method: %s | msg: %s"

type Entity interface {
	IsZeroValue() bool
}

type ABMController[T Entity] struct {
	name     string
	s        services.ABMService[T]
	Validate func(T) error
}

func NewABMController[T Entity](service services.ABMService[T]) ABMController[T] {
	return ABMController[T]{s: service}
}

func (abm *ABMController[Entity]) New(c *gin.Context) {

	log.Debugf(logTemplate, abm.name, "NEW", fmt.Sprintf("new request | body: %v", getBody(c)))

	var e Entity
	err := c.BindJSON(&e)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "NEW", err)
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	err = abm.Validate(e)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "NEW", err)
		ReturnError(c, http.StatusBadRequest, ValidationError, err.Error())
		return
	}

	e, err = abm.s.New(e)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "NEW", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	log.Debugf(logTemplate, abm.name, "NEW", fmt.Sprintf("success | response: %v", e))

	c.JSON(http.StatusCreated, e)
}

func (abm *ABMController[Entity]) Get(c *gin.Context) {

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, abm.name, "GET", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "GET", "cannot parse id: "+err.Error())
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := abm.s.Get(id)
	if err != nil {
		log.Errorf(logTemplate, abm.name, "GET", err)
		ReturnError(c, http.StatusInternalServerError, ServiceError, err.Error())
		return
	}

	if e.IsZeroValue() {
		log.Debugf(logTemplate, abm.name, "GET", "entity not found")
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}
	log.Debugf(logTemplate, abm.name, "NEW", fmt.Sprintf("success | response: %v", e))
	c.JSON(http.StatusOK, e)
}

func (abm *ABMController[Entity]) Edit(c *gin.Context) {

	log.Debugf(logTemplate, abm.name, "NEW", fmt.Sprintf("new request | body: %s", getBody(c)))

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, abm.name, "EDIT", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "EDIT", "cannot parse id: "+err.Error())
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	e, err := abm.s.Get(id)
	if err != nil {
		log.Errorf(logTemplate, abm.name, "GET", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}
	if e.IsZeroValue() {
		log.Debugf(logTemplate, abm.name, "EDIT", "entity not found")
		ReturnError(c, http.StatusNotFound, EntityNotFound, fmt.Sprintf("entity id '%d' not found", id))
		return
	}

	err = c.BindJSON(&e)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "EDIT", err)
		ReturnError(c, http.StatusBadRequest, EntityFormatError, err.Error())
		return
	}

	err = abm.Validate(e)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "EDIT", err)
		ReturnError(c, http.StatusBadRequest, ValidationError, err.Error())
		return
	}

	e, err = abm.s.Edit(id, e)
	if err != nil {
		log.Errorf(logTemplate, abm.name, "EDIT", err)
		ReturnError(c, http.StatusInternalServerError, RegisterError, err.Error())
		return
	}

	log.Debugf(logTemplate, abm.name, "NEW", fmt.Sprintf("success | response: %v", e))
	c.JSON(http.StatusCreated, e)

}

func (abm *ABMController[Entity]) Delete(c *gin.Context) {

	idStr, ok := c.Params.Get(IDParamName)
	if !ok || idStr == "" {
		log.Debugf(logTemplate, abm.name, "DELETE", "expected entity id")
		ReturnError(c, http.StatusBadRequest, MissingParams, "expected entity id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Debugf(logTemplate, abm.name, "DELETE", err)
		ReturnError(c, http.StatusBadRequest, MissingParams, "cannot parse id: "+err.Error())
		return
	}

	abm.s.Delete(id)
	log.Debugf(logTemplate, abm.name, "DELETE", "success")
	c.JSON(http.StatusNoContent, nil)
}

func getBody(c *gin.Context) string {

	// Crea un buffer para almacenar el cuerpo de la solicitud
	var requestBodyBuffer bytes.Buffer

	teeReader := io.TeeReader(c.Request.Body, &requestBodyBuffer)
	bodyBytes, err := io.ReadAll(teeReader)
	if err != nil {
		return "[error reading body]"
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return strings.ReplaceAll(string(bodyBytes), "\n", "")
}
