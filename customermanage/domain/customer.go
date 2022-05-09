package domain

type (
	// Customer type represent the registered customer.
	Customer struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	// CRUD inteface
	CustomerStore interface {
		Create(Customer) error
		Update(string, Customer) error
		Delete(string) error
		GetById(string) (Customer, error)
		GetAll() ([]Customer, error)
	}
)
