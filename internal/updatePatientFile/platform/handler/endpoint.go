package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updatePatientFile"
)

func MakeUpdatePatientFileEndpoint(u updatePatientFile.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePatientFileInternalRequest)
		resp, err := u.UpdatePatientFileSvc(req.ctx, req.IdPatient, req.IdPatientFile, req.StatePatient, req.HighDate, req.LowDate, req.Comorbidity, req.Symptom)
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
	IdPatient     string   `json:"idPatient"`
	IdPatientFile string   `json:"idPatientFile"`
	StatePatient  string   `json:"statePatient"`
	HighDate      string   `json:"highDate"`
	LowDate       string   `json:"lowDate"`
	Comorbidity   []string `json:"comorbidity"`
	Symptom       []string `json:"symptom"`
	ctx           context.Context
}
