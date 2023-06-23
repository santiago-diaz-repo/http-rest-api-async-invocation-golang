package handler

import (
	v1 "http-rest-api-versioning-golang/data/v1"
	v2 "http-rest-api-versioning-golang/data/v2"
	"net/http"
	"strings"

	"github.com/google/jsonapi"
)

// getAllCustomersByAcceptAccept returns version 1 and 2 of customers depends on the version sent by Accept header
func (h *Handler) getAllCustomersByAcceptAccept(w http.ResponseWriter, r *http.Request) {
	jsonapiRuntime := jsonapi.NewRuntime().Instrument("customers.list")

	if r.Method != http.MethodGet {
		http.Error(w, "unsuported method", http.StatusMethodNotAllowed)
		return
	}

	acceptHeader := r.Header.Get(headerAccept)
	headersAccept := strings.Split(acceptHeader, headerSeparator)
	if headersAccept[0] != jsonapi.MediaType {
		http.Error(w, "unsupported media Type", http.StatusUnsupportedMediaType)
		return
	}

	version := strings.Split(headersAccept[1], equal)

	switch version[1] {
	case "1":
		data := v1.NewLoansManagerV1()
		loans := data.GetLoans()

		w.Header().Set(headerContentType, acceptHeader)
		w.WriteHeader(http.StatusOK)
		if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "2":
		data := v2.NewLoansManagerV2()
		loans := data.GetLoans()

		w.Header().Set(headerContentType, acceptHeader)
		w.WriteHeader(http.StatusOK)
		if err := jsonapiRuntime.MarshalPayload(w, loans); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "unsupported version", http.StatusBadRequest)
		return
	}
}
