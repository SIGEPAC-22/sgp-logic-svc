package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/createConmorbility"
)

func MakeCreateConmorbilityEndpoint(c createConmorbility.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateConmorbilityInternalRequest)
		resp, err := c.CreateConmorbilitySvc(req.ctx, req.NameConmorbility, req.DescriptionConmorbility)
		return CreateConmorbilityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type CreateConmorbilityInternalResponse struct {
	Response interface{}
	Err      error
}

type CreateConmorbilityInternalRequest struct {
	NameConmorbility        string `json:"name_conmorbility"`
	DescriptionConmorbility string `json:"description_conmorbility"`
	ctx                     context.Context
}
