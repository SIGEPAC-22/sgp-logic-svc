package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/createSymptom"
)

func MakeCreateSymptomEndpoint(c createSymptom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateSymptomInternalRequest)
		resp, err := c.CreateSymptomSvc(req.ctx, req.NameSymptom, req.DescriptionSymptom)
		return CreateSymptomInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type CreateSymptomInternalResponse struct {
	Response interface{}
	Err      error
}

type CreateSymptomInternalRequest struct {
	NameSymptom        string `json:"nameSymptom"`
	DescriptionSymptom string `json:"descriptionSymptom"`
	ctx                context.Context
}
