package datastore

import (
	"hospi/model"

	"gofr.dev/pkg/gofr"
)

type Patient interface {
	// GetByID retrieves a patient record based on its ID.
	GetByID(ctx *gofr.Context, id int) (*model.Patient, error)
	// Create inserts a new patient record into the database.
	Create(ctx *gofr.Context, model *model.Patient) (*model.Patient, error)
	// Update updates an existing patient record with the provided information.
	Update(ctx *gofr.Context, model *model.Patient) (*model.Patient, error)
	// Delete removes a patient record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}