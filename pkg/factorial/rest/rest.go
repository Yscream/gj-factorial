package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Yscream/go-factorial/pkg/factorial/models"
)

// i decided to create respondJSON and respondError funcs because their logic is reused in the code
func RespondJSON(w http.ResponseWriter, code int, response interface{}) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func RespondError(w http.ResponseWriter, code int, errMsg string) {
	errResponse := &models.ErrorResponse{Error: errMsg}
	RespondJSON(w, code, errResponse)
}
