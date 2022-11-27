package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updateSymptom"
)

func MakeUpdateSymptomEndpoint(u updateSymptom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateSymptomInternalRequest)
		resp, err := u.UpdateSymptomSvc(req.ctx, req.Id, req.NameSymptom, req.DescriptionSymptom)
		return UpdateSymptomInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdateSymptomInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdateSymptomInternalRequest struct {
	Id                 string `json:"id"`
	NameSymptom        string `json:"nameSymptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
	ctx                context.Context
}
