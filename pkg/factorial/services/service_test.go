package services

import (
	"testing"

	"github.com/Yscream/go-factorial/pkg/factorial/mock"
	"github.com/stretchr/testify/assert"
	
	entities "github.com/Yscream/go-factorial"
)

func TestFactorialService_Calculate(t *testing.T) {
	mockCalculator := &mock.FactorialCalculatorMock{
		CalculateFunc: func(n int) int {
			result := 1
			for i := 2; i <= n; i++ {
				result *= i
			}

			return result
		},
	}

	factorialSvc := NewFactorialService(mockCalculator)

	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{8, 40320},
	}

	for _, tc := range testCases {
		result := factorialSvc.Calculate(tc.input)

		assert.Equal(t, tc.expected, result)
	}
}

func TestFactorialService_ConcurrentFactorialsCalculation(t *testing.T) {
	mockCalculator := &mock.FactorialCalculatorMock{
		CalculateFunc: func(n int) int {
			result := 1
			for i := 2; i <= n; i++ {
				result *= i
			}

			return result
		},
	}

	factorialSvc := NewFactorialService(mockCalculator)

	returnIntPointer := func(n int) *int { return &n }

	testCases := []struct {
		input    *entities.Numbers
		expected *entities.Numbers
	}{
		{
			&entities.Numbers{
				A: returnIntPointer(0),
				B: returnIntPointer(0),
			},
			&entities.Numbers{
				A: returnIntPointer(1),
				B: returnIntPointer(1),
			},
		},
		{
			&entities.Numbers{
				A: returnIntPointer(1),
				B: returnIntPointer(1),
			},
			&entities.Numbers{
				A: returnIntPointer(1),
				B: returnIntPointer(1),
			},
		},
		{
			&entities.Numbers{
				A: returnIntPointer(5),
				B: returnIntPointer(6),
			},
			&entities.Numbers{
				A: returnIntPointer(120),
				B: returnIntPointer(720),
			},
		},
		{
			&entities.Numbers{
				A: returnIntPointer(8),
				B: returnIntPointer(9),
			},
			&entities.Numbers{
				A: returnIntPointer(40320),
				B: returnIntPointer(362880),
			},
		},
		{
			&entities.Numbers{
				A: returnIntPointer(11),
				B: returnIntPointer(12),
			},
			&entities.Numbers{
				A: returnIntPointer(39916800),
				B: returnIntPointer(479001600),
			},
		},
	}

	for _, tc := range testCases {
		result := factorialSvc.CalculateConcurrently(tc.input)

		assert.Equal(t, *tc.expected.A, *result.A)
		assert.Equal(t, *tc.expected.B, *result.B)
	}
}
