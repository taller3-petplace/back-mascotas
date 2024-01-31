package domain

import (
	"petplace/back-mascotas/src/model"
	"time"
)

// VaccineResponse response from Treatments service
type VaccineResponse struct {
	ID   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func (r *VaccineResponse) ToModel() model.Vaccine {
	return model.Vaccine{
		ID:          0,
		Animal:      "unknown",
		Name:        r.Name,
		Description: "unknown",
		Scheduled:   0,
		AppliedAt:   &r.Date,
	}
}
