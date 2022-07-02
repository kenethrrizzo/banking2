package customers

import "errors"

type CustomerService interface {
	ListAllCustomers() ([]Customer, error)
	GetCustomerByID(uint) (*Customer, error)
	CreateNewCustomer(Customer) (*Customer, error)
	UpdateCustomer(Customer) (*Customer, error)
	DeleteCustomer(uint) error
}

type Service struct {
	repo CustomerRepository
}

func NewService(repo CustomerRepository) *Service {
	return &Service{repo}
}

func (s *Service) ListAllCustomers() ([]Customer, error) {
	return s.repo.GetAllCustomers()
}

func (s *Service) GetCustomerByID(id uint) (*Customer, error) {
	return s.repo.GetCustomerByID(id)
}

// Crea un nuevo cliente en la base de datos
func (s *Service) CreateNewCustomer(cust Customer) (*Customer, error) {
	if !cust.IsOfLegalAge() {
		return nil, errors.New("the customer must be of legal age")
	}
	return s.repo.CreateCustomer(cust)
}

func (s *Service) UpdateCustomer(cust Customer) (*Customer, error) {
	return s.repo.UpdateCustomer(cust)
}

func (s *Service) DeleteCustomer(id uint) error {
	return s.repo.DeleteCustomer(id)
}
