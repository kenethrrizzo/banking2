package customers

import "time"

type Customer struct {
	ID                   uint `gorm:"primaryKey;autoIncrement"`
	Name                 string
	Surname              string
	IdentificationType   byte
	IdentificationNumber string `gorm:"unique"`
	Country              string
	City                 string
	ZipCode              string
	DateOfBirth          time.Time
	Status               byte
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
