package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/updateComorbidity"
)

func MakeUpdateComorbidityEndpoint(u updateComorbidity.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateComorbidityInternalRequest)
		resp, err := u.UpdateComorbiditySvc(req.ctx, req.Id, req.NameComorbidity, req.DescriptionComorbidity)
		return UpdateComorbidityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type UpdateComorbidityInternalResponse struct {
	Response interface{}
	Err      error
}

type UpdateComorbidityInternalRequest struct {
	Id                     string `json:"id"`
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
	ctx                    context.Context
}
