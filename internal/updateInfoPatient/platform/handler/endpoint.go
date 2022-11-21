package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updateInfoPatient"
)

func MakeUpdateInfoPatientEndpoint(u updateInfoPatient.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateInfoPatientInternalRequest)
		//resp, err := u.UpdateInfoPatientSvc(req.ctx, req.Id, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DateBirth, req.DocumentType, req.DocumentNumber, req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department, req.Country, req.PatientFile, req.PatientSex)
		resp, err := u.UpdateInfoPatientSvc(req.ctx, req.Id, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DocumentType, req.DocumentNumber, req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department, req.Country)
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
	Country                      string `json:"country"`
	PatientFile                  string `json:"patientFile"`
	ctx                          context.Context
}
