package updateInfoPatient

import (
	"context"
)

type Repository interface {
	SelectInfoPatientRepo(ctx context.Context, id int) (SelectInfoPatientResponse, error)
	UpdateInfoPatientRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int) (bool, error)
}

type Service interface {
	UpdateInfoPatientSvc(ctx context.Context, Id string, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType string, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department string) (UpdateInfoPatientResponse, error)
}

type SelectInfoPatientResponse struct {
	DocumentType int `json:"documentType"`
	Department   int `json:"department"`
	Foreign      int `json:"foreign"`
	Pregnant     int `json:"pregnant"`
}

type UpdateInfoPatientRequest struct {
	Id                           string `json:"id"`
	FirstName                    string `json:"firstName"`
	SecondName                   string `json:"secondName"`
	LastFirstName                string `json:"lastFirstName"`
	LastSecondName               string `json:"lastSecondName"`
	DocumentType                 string `json:"documentType"`
	DocumentNumber               string `json:"documentNumber"`
	CellPhoneNumber              string `json:"cellPhoneNumber"`
	PhoneNumber                  string `json:"phoneNumber"`
	ResponsibleFamily            string `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber string `json:"responsibleFamilyPhoneNumber"`
	Department                   string `json:"department"`
}

type UpdateInfoPatientResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}
