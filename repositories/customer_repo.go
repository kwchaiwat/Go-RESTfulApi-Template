package repositories

type Customer struct {
	CustomerId  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	Status      bool   `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(id int) (*Customer, error)
}
