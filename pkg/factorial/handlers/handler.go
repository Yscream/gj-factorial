package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Yscream/go-factorial/pkg/factorial/services"
	"github.com/julienschmidt/httprouter"

	entities "github.com/Yscream/go-factorial"
)

type Handler struct {
	factorialSvc *services.FactorialService
}

func NewHandler(factorialSvc *services.FactorialService) *Handler {
	return &Handler{
		factorialSvc: factorialSvc,
	}
}

func (h *Handler) NewCalculateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	inputNumbers := entities.Numbers{}
	err = json.Unmarshal(b, &inputNumbers)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateData(&inputNumbers); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	factorials := h.factorialSvc.CalculateConcurrently(&inputNumbers)

	respondJSON(w, http.StatusOK, factorials)
}
