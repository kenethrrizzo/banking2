package customers

import (
	cusd "github.com/kenethrrizzo/banking2/domain/customers"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&Customer{})
	return &Store{db: db}
}

func (s Store) GetAllCustomers() ([]cusd.Customer, error) {
	var results []Customer
	if err := s.db.Find(&results).Error; err != nil {
		return nil, err
	}
	var custs = make([]cusd.Customer, len(results))
	for i, result := range results {
		custs[i] = *toDomainModel(&result)
	}
	return custs, nil
}

func (s Store) GetCustomerByID(id uint) (*cusd.Customer, error) {
	var result Customer
	if err := s.db.First(&result, id).Error; err != nil {
		return nil, err
	}
	return toDomainModel(&result), nil
}

func (s Store) CreateCustomer(cust cusd.Customer) (*cusd.Customer, error) {
	if err := s.db.Create(toDBModel(&cust)).Error; err != nil {
		return nil, err
	}
	return &cust, nil
}

func (s Store) UpdateCustomer(cust cusd.Customer) (*cusd.Customer, error) {
	if err := s.db.Save(toDBModel(&cust)).Error; err != nil {
		return nil, err
	}
	return &cust, nil
}

// Logic delete
func (s Store) DeleteCustomer(id uint) error {
	return s.db.Model(&Customer{}).Where(id).Update("Status", 0).Error
}
