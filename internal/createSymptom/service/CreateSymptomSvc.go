package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/createSymptom"
	"sgp-logic-svc/kit/constants"
)

type CreateSymptomSvc struct {
	repoDB createSymptom.Repository
	logger kitlog.Logger
}

func NewCreateSymptomSvc(repoDB createSymptom.Repository, logger kitlog.Logger) *CreateSymptomSvc {
	return &CreateSymptomSvc{repoDB: repoDB, logger: logger}
}

func (c CreateSymptomSvc) CreateSymptomSvc(ctx context.Context, NameSymptom string, DescriptionSymptom string) (createSymptom.CreateSymptomResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := c.repoDB.CreateSymptomRepo(ctx, NameSymptom, DescriptionSymptom)
	if err != nil {
		c.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createSymptom.CreateSymptomResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			c.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return createSymptom.CreateSymptomResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return createSymptom.CreateSymptomResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
