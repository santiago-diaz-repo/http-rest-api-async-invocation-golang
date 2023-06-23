package handler

import (
	v1 "http-rest-api-versioning-golang/data/v1"
	v2 "http-rest-api-versioning-golang/data/v2"
	"net/http"

	"github.com/google/jsonapi"
)

// getAllCustomersByUrlV1 returns version 1 of customers
func (h *Handler) getAllCustomersByUrlV1(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet {
		http.Error(w, "unsuported method", http.StatusMethodNotAllowed)
		return
	}

	data := v1.NewLoansManagerV1()
	loans := data.GetLoans()

	w.Header().Set(headerContentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getAllCustomersByUrlV2 returns version 2 of customers
func (h *Handler) getAllCustomersByUrlV2(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet {
		http.Error(w, "unsuported method", http.StatusMethodNotAllowed)
		return
	}

	data := v2.NewLoansManagerV2()
	loans := data.GetLoans()

	w.Header().Set(headerContentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
