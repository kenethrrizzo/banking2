package customers

type CustomerRepository interface {
	GetAllCustomers() ([]Customer, error)
	GetCustomerByID(uint) (*Customer, error)
	CreateCustomer(Customer) (*Customer, error)
	UpdateCustomer(Customer) (*Customer, error)
	DeleteCustomer(uint) error
}
