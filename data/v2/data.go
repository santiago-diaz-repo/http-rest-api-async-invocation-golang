package v2

import (
	"http-rest-api-versioning-golang/models/v2"
)

type Data struct {}

func NewData() Data {
	return Data{}
}

func (d *Data) GetLoans() []*v2.Customer {
	var customers []*v2.Customer
	customer1 := &v2.Customer{
		ID: 221234,
		FirstName: "Test1",
		LastName: "Test1",
		Loans: []*v2.Loan{
			{
				ID: 77565543,
				Balance: 343354.67,
			},
		},
	}

	customer2 := &v2.Customer{
		ID: 964532,
		FirstName: "Test2",
		LastName: "Test2",
		Loans: []*v2.Loan{
			{
				ID: 5466878,
				Balance: 343354.67,
			},
			{
				ID: 67654590,
				Balance: 100006.50,
			},
		},
	}
	customers = append(customers,customer1)
	customers = append(customers, customer2)
	return customers
}

