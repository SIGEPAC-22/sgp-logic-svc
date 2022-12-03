package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updatePatientFile"
)

func MakeUpdatePatientFileEndpoint(u updatePatientFile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePatientFileInternalRequest)
		resp, err := u.UpdatePatientFileSvc(req.ctx, req.IdPatient, req.IdPatientFile, updatePatientFile.StatePatient(req.StatePatient), req.HighDate, req.LowDate, updatePatientFile.Comorbidity(req.Comorbidity), updatePatientFile.Symptom(req.Symptom))
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
	IdPatient     string       `json:"idPatient"`
	IdPatientFile string       `json:"idPatientFile"`
	StatePatient  StatePatient `json:"statePatient"`
	HighDate      string       `json:"highDate"`
	LowDate       string       `json:"lowDate"`
	Comorbidity   Comorbidity  `json:"comorbidity"`
	Symptom       Symptom      `json:"symptom"`
	ctx           context.Context
}

type Comorbidity []struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Symptom []struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type StatePatient struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
