package main

import (
	"log"
	"net/http"

	"github.com/Yscream/go-factorial/pkg/factorial/router"
	"github.com/Yscream/go-factorial/pkg/factorial/services"
)

func main() {
	newFactorialCalculator := &services.FactorialService{}
	svc := services.NewFactorialService(newFactorialCalculator)
	r := router.New(svc)

	log.Println("server is starting on port :8989.....")
	log.Fatal(http.ListenAndServe(":8989", r))
}
