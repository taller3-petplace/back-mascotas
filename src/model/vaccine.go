package model

import "time"

type Vaccine struct {
	ID          uint       `json:"id"`
	Animal      AnimalType `json:"animal"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Scheduled   uint       `json:"scheduled,omitempty"`
	AppliedAt   *time.Time `json:"applied_at,omitempty"`
}

func (v Vaccine) IsZeroValue() bool {

	var zeroValue Vaccine

	result := v.ID == zeroValue.ID
	result = result && (v.Name == zeroValue.Name)
	result = result && (v.Animal == zeroValue.Animal)
	result = result && (v.Description == zeroValue.Description)
	result = result && (v.Scheduled == zeroValue.Scheduled)
	result = result && (v.AppliedAt == zeroValue.AppliedAt)

	return result
}

type VaccinationPlan struct {
	Name    string
	Type    string
	OwnerID string
	Applied []Vaccine
	Pending []Vaccine
}

// VaccineResponse response from Treatments service
type VaccineResponse struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (r VaccineResponse) ToModel() Vaccine {

	return Vaccine{
		ID:          0,
		Animal:      "unknown",
		Name:        r.Name,
		Description: "unknown",
		Scheduled:   0,
		AppliedAt:   &r.Date,
	}
}
