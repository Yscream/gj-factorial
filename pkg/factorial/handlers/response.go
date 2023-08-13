package handlers

import (
	"encoding/json"
	"net/http"

	entities "github.com/Yscream/go-factorial"
)

func respondJSON(w http.ResponseWriter, code int, response interface{}) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func respondError(w http.ResponseWriter, code int, errMsg string) {
	respondJSON(w, code, newError(errMsg))
}

func newError(errMsg string) *entities.ErrorResponse {
	return &entities.ErrorResponse{
		Error: errMsg,
	}
}