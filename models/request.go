package models

// Request for body and location header
type Request struct {
	Name string `json:"name"`
}

// RequestPush request for making a callback
type RequestPush struct {
	Name string `json:"name"`
	UrlCallback string `json:"urlCallback"`
}