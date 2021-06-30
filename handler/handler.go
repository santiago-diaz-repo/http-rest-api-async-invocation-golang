package handler

import "net/http"

type Handler struct {}

const (
	applicationJson = "application/json"
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
	headerSeparator =";"
	equal = "="
)

func (h *Handler) HandleRequest() {

	http.HandleFunc("/api/customers",h.getAllCustomersAccept)
	http.HandleFunc("/api/v1/customers",h.getAllCustomersUrlV1)
	http.HandleFunc("/api/v2/customers",h.getAllCustomersUrlV2)
	http.ListenAndServe(":8080",nil)

}
