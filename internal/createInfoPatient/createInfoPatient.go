package createInfoPatient

import (
	"context"
)

type Repository interface {
	CreateInfoPatientRepo(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (bool, error)
}

type Service interface {
	CreateInfoPatientSvc(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (CreateInfoPatientResponse, error)
}

type CreateInfoPatientRequest struct {
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

type CreateInfoPatientResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
