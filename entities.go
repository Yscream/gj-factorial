package entities

//go:generate moq -pkg mock -out pkg/factorial/mock/factorial_calculator_mock.go . FactorialCalculator

type Numbers struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type FactorialCalculator interface {
	Calculate(n int) int
	CalculateConcurrently(inputNumbers *Numbers) *Numbers
}
