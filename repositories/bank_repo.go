package repositories

type Bank struct {
	Id             int     `db:"id"`
	AccountNumber  string  `db:"account_number"`
	Trust          float64 `db:"trust"`
	TransactionFee int     `db:"transaction_fee"`
}

type BankRepository interface {
	GetAll() ([]Bank, error)
	GetById(id int) (*Bank, error)
}
