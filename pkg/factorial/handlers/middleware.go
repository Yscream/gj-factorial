package handlers

import (
	"errors"

	"github.com/Yscream/go-factorial/pkg/factorial/models"
)

// func checks model for empty fields and negative numbers
func validateModel(d *models.Data) error {
	switch {
	case d.A == nil:
		return errors.New("incorrect input")
	case d.B == nil:
		return errors.New("incorrect input")
	case *d.A < 0:
		return errors.New("incorrect input")
	case *d.B < 0:
		return errors.New("incorrect input")
	}

	return nil
}
