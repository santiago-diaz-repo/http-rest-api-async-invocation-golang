package handler

import (
	"github.com/google/jsonapi"
	v1 "http-rest-api-versioning-golang/data/v1"
	v2 "http-rest-api-versioning-golang/data/v2"
	"net/http"
	"strings"
)

// getAllCustomersAccept returns version 1 and 2 of customers depends on the version sent by Accept header
func (h *Handler) getAllCustomersAccept(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet{
		http.Error(w, "Only GET", http.StatusMethodNotAllowed)
		return
	}

	acceptHeader := r.Header.Get(headerAccept)
	headersAccept := strings.Split(acceptHeader,headerSeparator)
	if headersAccept[0] != jsonapi.MediaType {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	version := strings.Split(headersAccept[1],equal)

	switch version[1] {
	case "1":
		data := v1.NewData()
		loans := data.GetLoans()

		w.Header().Set(headerContentType, acceptHeader)
		w.WriteHeader(http.StatusOK)
		if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "2":
		data := v2.NewData()
		loans := data.GetLoans()

		w.Header().Set(headerContentType, acceptHeader)
		w.WriteHeader(http.StatusOK)
		if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Version is not supported", http.StatusBadRequest)
		return
	}
}