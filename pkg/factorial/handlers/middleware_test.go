package handlers

import (
	"testing"

	entities "github.com/Yscream/go-factorial"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ValidateData(t *testing.T) {
	returnIntPointer := func(n int) *int { return &n }

	testCases := []struct {
		name        string
		numbers     *entities.Numbers
		expectedErr bool
	}{
		{
			name: "ValidInput",
			numbers: &entities.Numbers{
				A: returnIntPointer(5),
				B: returnIntPointer(6),
			},
			expectedErr: false,
		},
		{
			name: "NilInput",
			numbers: &entities.Numbers{
				A: nil,
				B: nil,
			},
			expectedErr: true,
		},
		{
			name: "NegativInput",
			numbers: &entities.Numbers{
				A: returnIntPointer(-5),
				B: returnIntPointer(-6),
			},
			expectedErr: true,
		},
		{
			name: "NilA",
			numbers: &entities.Numbers{
				A: nil,
				B: returnIntPointer(3),
			},
			expectedErr: true,
		},
		{
			name: "NilB",
			numbers: &entities.Numbers{
				A: returnIntPointer(2),
				B: nil,
			},
			expectedErr: true,
		},
		{
			name: "NegativeA",
			numbers: &entities.Numbers{
				A: returnIntPointer(-5),
				B: returnIntPointer(9),
			},
			expectedErr: true,
		},
		{
			name: "NegativeB",
			numbers: &entities.Numbers{
				A: returnIntPointer(3),
				B: returnIntPointer(-5),
			},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateData(tc.numbers)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
