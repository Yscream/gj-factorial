package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yscream/go-factorial/pkg/factorial/mock"
	"github.com/Yscream/go-factorial/pkg/factorial/services"
	"github.com/stretchr/testify/assert"

	entities "github.com/Yscream/go-factorial"
)

func TestHandler_NewCalculateHandler(t *testing.T) {
	mockCalculator := &mock.FactorialCalculatorMock{
		CalculateFunc: func(n int) int {
			result := 1
			for i := 2; i <= n; i++ {
				result *= i
			}
			return result
		},
	}

	mockService := services.NewFactorialService(mockCalculator)
	handler := NewHandler(mockService)

	returnIntPointer := func(n int) *int { return &n }

	testCases := []struct {
		name            string
		inputBody       string
		expectedStatus  int
		expectedNumbers *entities.Numbers
	}{
		{
			name:           "ValidInput",
			inputBody:      `{"A": 5, "B": 3}`,
			expectedStatus: http.StatusOK,
			expectedNumbers: &entities.Numbers{
				A: returnIntPointer(120),
				B: returnIntPointer(6),
			},
		},
		{
			name:           "InvalidInput",
			inputBody:      `{"A": -5, "B": 3}`,
			expectedStatus: http.StatusBadRequest,
			expectedNumbers: &entities.Numbers{
				A: new(int),
				B: new(int),
			},
		},
		{
			name:           "EmptyInput",
			inputBody:      "",
			expectedStatus: http.StatusBadRequest,
			expectedNumbers: &entities.Numbers{
				A: new(int),
				B: new(int),
			},
		},
		{
			name:           "InvalidJSON",
			inputBody:      `invalid json`,
			expectedStatus: http.StatusBadRequest,
			expectedNumbers: &entities.Numbers{
				A: new(int),
				B: new(int),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/calculate", bytes.NewBufferString(tc.inputBody))

			handler.NewCalculateHandler(rec, req, nil)

			resp := rec.Result()
			defer resp.Body.Close()

			if ok := assert.Equal(t, tc.expectedStatus, resp.StatusCode); ok {
				b, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("could not read body, %x; err: %v", b, err)
				}
				defer req.Body.Close()

				responseNumbers := entities.Numbers{}
				err = json.Unmarshal(b, &responseNumbers)
				if err != nil {
					respondError(rec, http.StatusBadRequest, err.Error())
					return
				}

				assert.Equal(t, *tc.expectedNumbers.A, *responseNumbers.A)
				assert.Equal(t, *tc.expectedNumbers.B, *responseNumbers.B)
			}
		})
	}
}
