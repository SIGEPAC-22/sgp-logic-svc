package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updatePatientFile"
)

func MakeUpdatePatientFileEndpoint(u updatePatientFile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePatientFileInternalRequest)
		resp, err := u.UpdatePatientFileSvc(req.ctx, req.Id, req.FirstName, req.SecondName, req.LastFirstName, req.LastSecondName, req.DocumentType, req.DocumentNumber, req.CellPhoneNumber, req.PhoneNumber, req.ResponsibleFamily, req.ResponsibleFamilyPhoneNumber, req.Department)
		return UpdatePatientFileInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdatePatientFileInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdatePatientFileInternalRequest struct {
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
	PatientFile                  string `json:"patientFile"`
	ctx                          context.Context
}
