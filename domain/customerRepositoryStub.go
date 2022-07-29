package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "Jakarta", "12345", "2022-01-02", "1"},
		{"2", "User2", "Surabaya", "67890", "2022-06-03", "1"},
		{"3", "User4", "Bekasi", "13544", "2022-02-02", "1"},
		{"4", "User5", "Tangerang", "12346", "2022-05-02", "1"},
	}

	return CustomerRepositoryStub{Customer: customers}
}
