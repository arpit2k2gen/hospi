package datastore

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"hospi/model"
)

type hospi struct{}

func New() *hospi {
	return &hospi{}
}

func (s *hospi) GetByID(ctx *gofr.Context, id int) (*model.Patient, error) {
	var resp model.Patient
	strId := strconv.Itoa(id)
	err := ctx.DB().QueryRowContext(ctx, " SELECT PatientID, name, gender, roomNumber, diagnosis FROM Patients where PatientID=?", strId).
		Scan(&resp.PatientID, &resp.Name, &resp.Gender, &resp.RoomNumber, &resp.Diagnosis)
	switch err {
	case sql.ErrNoRows:
		strId := strconv.Itoa(id)
		return &model.Patient{}, errors.EntityNotFound{Entity: "Patients", ID: strId}
	case nil:
		return &resp, nil
	default:
		return &model.Patient{}, err
	}
}

func (s *hospi) Create(ctx *gofr.Context, patient *model.Patient) (*model.Patient, error) {
	var resp model.Patient

	res, err := ctx.DB().ExecContext(ctx, "INSERT INTO Patients (name, gender, roomNumber, diagnosis) VALUES (?,?,?,?)", patient.Name, patient.Gender, patient.RoomNumber, patient.Diagnosis)

	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}
	err = ctx.DB().QueryRowContext(ctx, "SELECT * FROM Patients WHERE PatientID = ?", id).Scan(&resp.PatientID, &resp.Name, &resp.Gender, &resp.RoomNumber, &resp.Diagnosis)
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}

	return &resp, nil
}

func (s *hospi) Update(ctx *gofr.Context, patient *model.Patient) (*model.Patient, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE Patients SET name=?, gender=?, roomNumber=?, diagnosis=? WHERE PatientID=?",
		patient.Name, patient.Gender, patient.RoomNumber, patient.Diagnosis, patient.PatientID)
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}

	return patient, nil
}

func (s *hospi) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM Patients where PatientID=?", id)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
