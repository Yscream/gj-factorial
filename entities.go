package entities

type Data struct {
	A *int `json:"a"`
	B *int `json:"b"`
}

type FactorialCalculator interface {
	CalculateFactorials(d *Data) *Data
}

type ErrorResponse struct {
	Error string `json:"error"`
}