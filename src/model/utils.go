package model

import (
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(date string) (Date, error) {

	birthDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: birthDate}, nil
}

// Implementa la interfaz Unmarshaler para MiFecha
func (f *Date) UnmarshalJSON(data []byte) error {

	// Intenta deserializar la cadena como una fecha en el formato "aaaa-mm-dd"
	dateString := strings.Trim(string(data), `"`)
	parsedTime, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		parsedTime, err = time.Parse(time.RFC3339, dateString)
		if err != nil {
			return err
		}
	}

	// Asigna la fecha deserializada al campo MiFecha
	f.Time = parsedTime
	return nil
}
