package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/deleteSymptom"
)

func MakeDeleteSymptomEndpoint(d deleteSymptom.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteSymptomInternalRequest)
		resp, err := d.DeleteSymptomSvc(req.ctx, req.Id)
		return DeleteSymptomInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type DeleteSymptomInternalResponse struct {
	Response interface{}
	Err      error
}

type DeleteSymptomInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
