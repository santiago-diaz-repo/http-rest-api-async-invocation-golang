package v2

import (
	"fmt"
	"github.com/google/jsonapi"
)

const host = "http://localhost:8080/customers"

type Customer struct {
	ID            int     `jsonapi:"primary,customers"`
	FirstName     string  `jsonapi:"attr,firstName"`
	LastName      string  `jsonapi:"attr,lastName"`
	Loans         []*Loan `jsonapi:"relation,loans"`
}

// JSONAPILinks builds self link
func (c Customer) JSONAPILinks() *jsonapi.Links {
	return &jsonapi.Links{
		"self": fmt.Sprintf("%s/%d",host, c.ID),
	}
}

// JSONAPIRelationshipLinks builds links of relationships
func (c Customer) JSONAPIRelationshipLinks(relation string) *jsonapi.Links {
	if relation == "loans" {
		return &jsonapi.Links{
			"related": fmt.Sprintf("%s/%d/loans",host, c.ID),
		}
	}
	return nil
}

// JSONAPIMeta build meta object
func (c Customer) JSONAPIMeta() *jsonapi.Meta {
	return &jsonapi.Meta{
		"detail": "This is version 2",
	}
}
