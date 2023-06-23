package handler

import (
	"fmt"
	"net/http"
)

const (
	applicationJson   = "application/json"
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
	headerSeparator   = ";"
	equal             = "="
)

type Handler struct{}

func (h *Handler) HandleRequest() {
	http.HandleFunc("/api/customers", h.getAllCustomersByAcceptAccept)
	http.HandleFunc("/api/v1/customers", h.getAllCustomersByUrlV1)
	http.HandleFunc("/api/v2/customers", h.getAllCustomersByUrlV2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("service cannot start: " + err.Error())
	}
	fmt.Println("service ready to receive requests...")
}
