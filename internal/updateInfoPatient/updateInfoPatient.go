package updateInfoPatient

import (
	"context"
)

type Repository interface {
	UpdateInfoPatientRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (bool, error)
}

type Service interface {
	UpdateInfoPatientSvc(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (UpdateInfoPatientResponse, error)
}

type UpdateInfoPatientRequest struct {
	Id                           int    `json:"id"`
	FirstName                    string `json:"firstName"`
	SecondName                   string `json:"secondName"`
	LastFirstName                string `json:"lastFirstName"`
	LastSecondName               string `json:"lastSecondName"`
	DateBirth                    string `json:"dateBirth"`
	DocumentType                 int    `json:"documentType"`
	DocumentNumber               string `json:"documentNumber"`
	CellPhoneNumber              string `json:"cellPhoneNumber"`
	PhoneNumber                  string `json:"phoneNumber"`
	ResponsibleFamily            string `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber string `json:"responsibleFamilyPhoneNumber"`
	Department                   int    `json:"department"`
	Country                      int    `json:"country"`
	PatientFile                  int    `json:"patientFile"`
	PatientSex                   int    `json:"patientSex"`
}

type UpdateInfoPatientResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
