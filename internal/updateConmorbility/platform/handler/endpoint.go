package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updateConmorbility"
)

func MakeUpdateUpdateConmorbilityEndpoint(u updateConmorbility.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateConmorbilityInternalRequest)
		resp, err := u.UpdateConmorbilitySvc(req.ctx, req.Id, req.NameConmorbility, req.DescriptionConmorbility)
		return UpdateConmorbilityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdateConmorbilityInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdateConmorbilityInternalRequest struct {
	Id                      string `json:"id"`
	NameConmorbility        string `json:"name_conmorbility"`
	DescriptionConmorbility string `json:"descripcion_conmorbility"`
	ctx                     context.Context
}
