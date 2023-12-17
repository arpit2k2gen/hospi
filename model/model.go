package model

type Patient struct {
	PatientID  int    `json:"PatientID"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	RoomNumber int    `json:"roomNumber"`
	Diagnosis  string `json:"diagnosis"`
}