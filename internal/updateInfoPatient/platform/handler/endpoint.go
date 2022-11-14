package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updateInfoPatient"
)

func MakeUpdateInfoPatientEndpoint(u updateInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateInfoPatientInternalRequest)
		resp, err := u.UpdateInfoPatientSvc(req.ctx, req.Id, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DateBirth, req.DocumentType, req.DocumentNumber, req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department, req.Country, req.PatientFile, req.PatientSex)
		return UpdateInfoPatientInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdateInfoPatientInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdateInfoPatientInternalRequest struct {
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
	ctx                          context.Context
}
