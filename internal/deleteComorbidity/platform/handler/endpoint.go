package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"sgp-logic-svc/internal/deleteComorbidity"
)

func MakeDeleteComorbidityEndpoint(d deleteComorbidity.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteComorbidityInternalRequest)
		resp, err := d.DeleteComorbiditySvc(req.ctx, req.Id)
		return DeleteComorbidityInternalResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type DeleteComorbidityInternalResponse struct {
	Response interface{}
	Err      error
}

type DeleteComorbidityInternalRequest struct {
	Id  string `json:"id"`
	ctx context.Context
}
