package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type errorResponse struct {
	Error string `json:"error"`
}

type data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func calculateFactorial(d *data) {
	//create two channels to pass the factorial results for A and B
	ch1 := make(chan int)
	ch2 := make(chan int)

	//run goroutines to calculate factorials for A and B
	go factorial(d.A, ch1)
	go factorial(d.B, ch2)

	//get results from channels and store them in structure d
	d.A = <-ch1
	d.B = <-ch2
}

func factorial(n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
}

//func checks model for empty fields and negative numbers
func validate(d *data) error {
	if d.A <= 0 || d.B <= 0 {
		return errors.New("incorrect input")
	}
	return nil
}

// i decided to create respondJSON and respondError funcs because their logic is reused in the code
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
	errResponse := errorResponse{Error: errMsg}
	respondJSON(w, code, errResponse)
}

func calculateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// read request body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// initialize model
	d := data{}
	// unmarshal bytes to the model
	err = json.Unmarshal(b, &d)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// check unmarshaled json for empty fields and negative numbers
	if err := validate(&d); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// call the calculate func
	calculateFactorial(&d)
	// send modified model back
	respondJSON(w, http.StatusOK, d)

}

func main() {
	router := httprouter.New()
	router.POST("/calculate", calculateHandler)

	log.Println("server is starting on port :8989.....")
	log.Fatal(http.ListenAndServe(":8989", router))
}
