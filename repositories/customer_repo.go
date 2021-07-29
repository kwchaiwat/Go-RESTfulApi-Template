package repositories

type Customer struct {
	Id             int     `db:"id"`
	AccountNumber  string  `db:"account_number"`
	Trust          float64 `db:"trust"`
	TransactionFee int     `db:"transaction_fee"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(id int) (*Customer, error)
}
