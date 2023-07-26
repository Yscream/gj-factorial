package services

import "github.com/Yscream/go-factorial/pkg/factorial/models"

type FactorialCalculator interface {
	Calculate(d *models.Data)
}

type FactorialService struct {
	FactorialCalculatorSvc FactorialCalculator
}

func NewFactorialService(fc FactorialCalculator) *FactorialService {
	return &FactorialService{
		FactorialCalculatorSvc: fc,
	}
}

func (s *FactorialService) Calculate(d *models.Data) {
	// create two channels to pass the factorial results for A and B
	ch1 := make(chan *int)
	ch2 := make(chan *int)

	// run goroutines to calculate factorials for A and B
	go factorial(d.A, ch1)
	go factorial(d.B, ch2)

	// get results from channels and store them in structure d
	d.A = <-ch1
	d.B = <-ch2
}

func factorial(n *int, ch chan<- *int) {
	result := 1
	for i := 2; i <= *n; i++ {
		result *= i
	}

	ch <- &result
}
