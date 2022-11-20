package createInfoPatient

import (
	"context"
)

type Repository interface {
	CreateInfoPatientRepo(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, patientSex int, foreign int) (bool, error)
	SelectInfoPatient(ctx context.Context, firstName string, secondName string, lastFirstName string, documentNumber string) (int, error)
	CreatePatientFileRepo(ctx context.Context, id int, pregnant bool) (bool, error)
}

type Service interface {
	CreateInfoPatientSvc(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType string, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department string, patientSex string, pregnant string, foreign string) (CreateInfoPatientResponse, error)
}

type CreateInfoPatientRequest struct {
	FirstName                    string `json:"firstName"`
	SecondName                   string `json:"secondName"`
	LastFirstName                string `json:"lastFirstName"`
	LastSecondName               string `json:"lastSecondName"`
	DateBirth                    string `json:"dateBirth"`
	DocumentType                 string `json:"documentType"`
	DocumentNumber               string `json:"documentNumber"`
	CellPhoneNumber              string `json:"cellPhoneNumber"`
	PhoneNumber                  string `json:"phoneNumber"`
	ResponsibleFamily            string `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber string `json:"responsibleFamilyPhoneNumber"`
	Department                   string `json:"department"`
	PatientSex                   string `json:"patientSex"`
	Pregnant                     bool   `json:"pregnant"`
	Foreign                      string `json:"foreign"`
}

type CreateInfoPatientResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}

type SelectInfoPatientResponse struct {
	Id int `json:"id"`
}
