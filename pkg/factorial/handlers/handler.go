package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Yscream/go-factorial/pkg/factorial/models"
	"github.com/Yscream/go-factorial/pkg/factorial/rest"
	"github.com/Yscream/go-factorial/pkg/factorial/services"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	svc *services.FactorialService
}

func NewHandler(svc *services.FactorialService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) NewCalculateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// read request body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		rest.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// initialize model
	d := models.Data{}
	// unmarshal bytes to the model
	err = json.Unmarshal(b, &d)
	if err != nil {
		rest.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check unmarshaled json for empty fields and negative numbers
	if err := validateModel(&d); err != nil {
		rest.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// call the calculate func
	h.svc.Calculate(&d)
	// send modified model back
	rest.RespondJSON(w, http.StatusOK, d)

}
