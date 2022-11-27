package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/createComorbidity"
)

func MakeCreateComorbidityEndpoint(c createComorbidity.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateComorbidityInternalRequest)
		resp, err := c.CreateComorbiditySvc(req.ctx, req.NameComorbidity, req.DescriptionComorbidity)
		return CreateComorbidityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type CreateComorbidityInternalResponse struct {
	Response interface{}
	Err      error
}

type CreateComorbidityInternalRequest struct {
	NameComorbidity        string `json:"nameComorbidity"`
	DescriptionComorbidity string `json:"descriptionComorbidity"`
	ctx                    context.Context
}
