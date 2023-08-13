package router

import (
	"github.com/Yscream/go-factorial/pkg/factorial/handlers"
	"github.com/Yscream/go-factorial/pkg/factorial/services"
	"github.com/julienschmidt/httprouter"
)

func New(svc *services.FactorialService) *httprouter.Router {
	handler := handlers.NewHandler(svc)

	router := httprouter.New()
	router.POST("/calculate", handler.NewCalculateHandler)

	return router
}
