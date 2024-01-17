package model

type Vaccine struct {
	ID          int        `json:"id"`
	Animal      AnimalType `json:"animal"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Scheduled   uint       `json:"scheduled"`
}

func (v Vaccine) IsZeroValue() bool {

	var zeroValue Vaccine

	result := v.ID == zeroValue.ID
	result = result && (v.Name == zeroValue.Name)
	result = result && (v.Animal == zeroValue.Animal)
	result = result && (v.Description == zeroValue.Description)
	result = result && (v.Scheduled == zeroValue.Scheduled)

	return result
}
