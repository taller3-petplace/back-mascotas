package pet_errors

import "errors"

var (
	EntityFormatError = errors.New("entity could not be mapped")
	RegisterError     = errors.New("error trying to register pet")
)
