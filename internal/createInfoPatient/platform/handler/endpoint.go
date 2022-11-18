package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/createInfoPatient"
)

func MakeCreateInfoPatientEndpoint(c createInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateInfoPatientInternalRequest)
		resp, err := c.CreateInfoPatientSvc(req.ctx, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DateBirth, req.DocumentType, req.DocumentNumber, req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department, req.PatientSex, req.Pregnant)
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
	DocumentType                 int    `json:"documentType"`
	DocumentNumber               string `json:"documentNumber"`
	CellPhoneNumber              string `json:"cellPhoneNumber"`
	PhoneNumber                  string `json:"phoneNumber"`
	ResponsibleFamily            string `json:"responsibleFamily"`
	ResponsibleFamilyPhoneNumber string `json:"responsibleFamilyPhoneNumber"`
	Department                   int    `json:"department"`
	PatientSex                   int    `json:"patientSex"`
	Pregnant                     bool   `json:"pregnant"`
	ctx                          context.Context
}
