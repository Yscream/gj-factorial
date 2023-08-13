package handlers

import (
	"errors"

	entities "github.com/Yscream/go-factorial"
)

// func checks model for empty fields and negative numbers
func validateData(numbers *entities.Numbers) error {
	switch {
	case numbers.A == nil:
		return errors.New("incorrect input")
	case numbers.B == nil:
		return errors.New("incorrect input")
	case *numbers.A < 0:
		return errors.New("incorrect input")
	case *numbers.B < 0:
		return errors.New("incorrect input")
	}

	return nil
}
