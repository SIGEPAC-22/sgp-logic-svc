package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/deleteConmorbility"
)

func MakeDeleteConmorbilityEndpoint(d deleteConmorbility.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteConmorbilityInternalRequest)
		resp, err := d.DeleteConmorbilitySvc(req.ctx, req.Id)
		return DeleteConmorbilityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type DeleteConmorbilityInternalResponse struct {
	Response interface{}
	Err      error
}

type DeleteConmorbilityInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
