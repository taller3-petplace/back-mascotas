package objects

import "time"

type Application struct {
	ID        uint `gorm:"primaryKey;autoIncrement;unique"`
	PetID     uint
	VaccineID uint
	AppliedAt time.Time
}
