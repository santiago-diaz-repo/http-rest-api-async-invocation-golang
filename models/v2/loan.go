package v2

type Loan struct {
	ID       int        `jsonapi:"primary,loans"`
	Balance    float32     `jsonapi:"attr,balance"`
}


