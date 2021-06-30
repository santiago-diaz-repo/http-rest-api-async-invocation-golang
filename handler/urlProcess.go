package handler

import (
	"github.com/google/jsonapi"
	v1 "http-rest-api-versioning-golang/data/v1"
	v2 "http-rest-api-versioning-golang/data/v2"
	"net/http"
)

// getAllCustomersUrlV1 returns version 1 of customers
func (h *Handler) getAllCustomersUrlV1(w http.ResponseWriter, r *http.Request){
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet{
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}

	data := v1.NewData()
	loans := data.GetLoans()

	w.Header().Set(headerContentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getAllCustomersUrlV2 returns version 2 of customers
func (h *Handler) getAllCustomersUrlV2(w http.ResponseWriter, r *http.Request){
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet{
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}

	data := v2.NewData()
	loans := data.GetLoans()

	w.Header().Set(headerContentType, applicationJson)
	w.WriteHeader(http.StatusOK)
	if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
