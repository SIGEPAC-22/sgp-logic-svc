package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/createInfoPatient"
)

func MakeCreateInfoPatientEndpoint(c createInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateInfoPatientInternalRequest)
		resp, err := c.CreateInfoPatientSvc(req.ctx, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DateBirth, req.DocumentType, req.DocumentNumber,
			req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department, req.PatientSex, req.Pregnant, req.Foreign)
		return CreateInfoPatientInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type CreateInfoPatientInternalResponse struct {
	Response interface{}
	Err      error
}

type CreateInfoPatientInternalRequest struct {
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
	Pregnant                     string `json:"pregnant"`
	Foreign                      string `json:"foreign"`
	ctx                          context.Context
}
