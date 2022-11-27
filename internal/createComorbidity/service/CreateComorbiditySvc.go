package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/createComorbidity"
	"sgp-logic-svc/kit/constants"
)

type CreateComorbiditySvc struct {
	repoDB createComorbidity.Repository
	logger kitlog.Logger
}

func NewCreateComorbiditySvc(repoDB createComorbidity.Repository, logger kitlog.Logger) *CreateComorbiditySvc {
	return &CreateComorbiditySvc{repoDB: repoDB, logger: logger}
}

func (c CreateComorbiditySvc) CreateComorbiditySvc(ctx context.Context, NameComorbidity string, DescriptionComorbidity string) (createComorbidity.CreateComorbidityResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := c.repoDB.CreateComorbidityRepo(ctx, NameComorbidity, DescriptionComorbidity)
	if err != nil {
		c.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createComorbidity.CreateComorbidityResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			c.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return createComorbidity.CreateComorbidityResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return createComorbidity.CreateComorbidityResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
