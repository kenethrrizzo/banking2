package customers

import "github.com/kenethrrizzo/banking2/domain/customers"

func toDomainModel(c *Customer) *customers.Customer {
	return &customers.Customer{
		ID:                   c.ID,
		Name:                 c.Name,
		Surname:              c.Surname,
		IdentificationType:   c.IdentificationType,
		IdentificationNumber: c.IdentificationNumber,
		Country:              c.Country,
		City:                 c.City,
		ZipCode:              c.ZipCode,
		DateOfBirth:          c.DateOfBirth,
		Status:               c.Status,
		CreatedAt:            c.CreatedAt,
		UpdatedAt:            c.UpdatedAt,
	}
}

func toDBModel(c *customers.Customer) *Customer {
	return &Customer{
		ID:                   c.ID,
		Name:                 c.Name,
		Surname:              c.Surname,
		IdentificationType:   c.IdentificationType,
		IdentificationNumber: c.IdentificationNumber,
		Country:              c.Country,
		City:                 c.City,
		ZipCode:              c.ZipCode,
		DateOfBirth:          c.DateOfBirth,
		Status:               c.Status,
		CreatedAt:            c.CreatedAt,
		UpdatedAt:            c.UpdatedAt,
	}
}
