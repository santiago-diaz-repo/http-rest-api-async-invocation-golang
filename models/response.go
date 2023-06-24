package models

// Response with uuid
type Response struct {
	Uuid string `json:"uuid""`
}

// ResponseAsync response with status of execution
type ResponseAsync struct {
	State bool `json:"state"`
}