package v1

type Loan struct {
	ID       int        `jsonapi:"primary,loans"`
	NickName   string        `jsonapi:"attr,nickName"`
	Balance    float32     `jsonapi:"attr,balance"`
}


