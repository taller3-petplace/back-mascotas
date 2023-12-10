package pet_errors

import "errors"

var (
	EntityFormatError = errors.New("entity could not be mapped")
	RegisterError     = errors.New("error trying to register pet")
	InvalidAnimalType = errors.New("invalid animal type")
	InvalidBirthDate  = errors.New("invalid birth_date")
	MissingParams     = errors.New("missing params on request")
	PetNotFound       = errors.New("pet not found")
	ServiceError      = errors.New("service error")
)
