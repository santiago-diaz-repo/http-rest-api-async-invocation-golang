package v1

import (
	"http-rest-api-versioning-golang/models/v1"
)

type Data struct {}

func NewData() Data {
	return Data{}
}

func (d *Data) GetLoans() []*v1.Customer {
	var customers []*v1.Customer
	customer1 := &v1.Customer{
		ID: 221234,
		FirstName: "Test1",
		LastName: "Test1",
		Address: "123 Test street",
		Loans: []*v1.Loan{
			{
				ID: 77565543,
				NickName: "loan-test1",
				Balance: 343354.67,
			},
		},
	}

	customer2 := &v1.Customer{
		ID: 964532,
		FirstName: "Test2",
		LastName: "Test2",
		Address: "987 Test2 street",
		Loans: []*v1.Loan{
			{
				ID: 5466878,
				NickName: "loan-test2-1",
				Balance: 343354.67,
			},
			{
				ID: 67654590,
				NickName: "loan-test2-2",
				Balance: 100006.50,
			},
		},
	}
	customers = append(customers,customer1)
	customers = append(customers, customer2)
	return customers
}

