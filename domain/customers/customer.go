package customers

import "time"

type Customer struct {
	ID                   uint
	Name                 string
	Surname              string
	IdentificationType   byte
	IdentificationNumber string
	Country              string
	City                 string
	ZipCode              string
	DateOfBirth          time.Time
	Status               byte
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (c Customer) IsOfLegalAge() bool {
	return c.DateOfBirth.AddDate(18, 0, 0).Before(time.Now())
}

func (c Customer) IsActive() bool {
	return c.Status == 1
}

func (c Customer) StatusAsText() string {
	if c.Status == 1 {
		return "Active"
	}
	return "Inactive"
}
