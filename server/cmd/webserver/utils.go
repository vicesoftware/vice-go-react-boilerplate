package main

import (
	"encoding/json"
	"net/http"

	"github.com/vicesoftware/vice-go-boilerplate/cmd/webserver/models"
	"github.com/vicesoftware/vice-go-boilerplate/pkg/database"
)

func Ok(w http.ResponseWriter, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func handleErrors(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		err := f(w, r)
		if err != nil {
			handleError(w, err)
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	if isInvalidRequest(err) || database.IsInvalidRequest(err) {
		http.Error(w, errToJSON(err), 400)
	} else if isNotFound(err) || database.IsNotFound(err) {
		http.Error(w, errToJSON(err), 404)
	} else {
		http.Error(w, errToJSON(err), 500)
	}
}

func isInvalidRequest(err error) bool {
	_, ok := err.(*invalidRequest)
	return ok
}

func isNotFound(err error) bool {
	_, ok := err.(*notFound)
	return ok
}

func errToJSON(err error) string {
	b, _ := json.Marshal(models.ErrorResponse{Error: err.Error()})
	return string(b)
}
