package router

import (
	"github.com/Yscream/go-factorial/pkg/factorial/handlers"
	"github.com/Yscream/go-factorial/pkg/factorial/services"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(svc *services.FactorialService) *httprouter.Router {
	h := handlers.NewHandler(svc)

	router := httprouter.New()
	router.POST("/calculate", h.NewCalculateHandler)

	return router
}
