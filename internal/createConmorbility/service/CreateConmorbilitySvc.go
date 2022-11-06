package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/createConmorbility"
	"sgp-logic-svc/kit/constants"
)

type CreateConmorbilitySvc struct {
	repoDB createConmorbility.Repository
	logger kitlog.Logger
}

func NewCreateConmorbilitySvc(repoDB createConmorbility.Repository, logger kitlog.Logger) *CreateConmorbilitySvc {
	return &CreateConmorbilitySvc{repoDB: repoDB, logger: logger}
}

func (c CreateConmorbilitySvc) CreateConmorbilitySvc(ctx context.Context, NameConmorbility string, DescriptionConmorbility string) (createConmorbility.CreateConmorbilityResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := c.repoDB.CreateConmorbilityRepo(ctx, NameConmorbility, DescriptionConmorbility)
	if err != nil {
		c.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createConmorbility.CreateConmorbilityResponse{
			ResponseCode: http.StatusBadRequest,
			Messagge:     "failed",
		}, constants.ErrorDataError

		if resp == false {
			c.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return createConmorbility.CreateConmorbilityResponse{
				ResponseCode: http.StatusBadRequest,
				Messagge:     "failed",
			}, constants.ErrorDataError
		}
	}
	return createConmorbility.CreateConmorbilityResponse{
		ResponseCode: http.StatusOK,
		Messagge:     "Successful",
	}, nil
}
