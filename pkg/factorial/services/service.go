package services

import (
	"sync"

	entities "github.com/Yscream/go-factorial"
)

type FactorialService struct {
	factorialCalculator entities.FactorialCalculator
}

func NewFactorialService(fc entities.FactorialCalculator) *FactorialService {
	return &FactorialService{
		factorialCalculator: fc,
	}
}

func (s *FactorialService) Calculate(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func (s *FactorialService) CalculateConcurrently(inputNumbers *entities.Numbers) *entities.Numbers {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		*inputNumbers.A = s.factorialCalculator.Calculate(*inputNumbers.A)
	}()

	go func() {
		defer wg.Done()
		*inputNumbers.B = s.factorialCalculator.Calculate(*inputNumbers.B)
	}()

	wg.Wait()

	return inputNumbers
}
