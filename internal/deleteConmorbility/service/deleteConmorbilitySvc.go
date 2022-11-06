package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/deleteConmorbility"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type DeleteConmorbilityService struct {
	repoDB deleteConmorbility.Repository
	logger kitlog.Logger
}

func NewDeleteConmorbilityService(repoDB deleteConmorbility.Repository, logger kitlog.Logger) *DeleteConmorbilityService {
	return &DeleteConmorbilityService{repoDB: repoDB, logger: logger}
}

func (d DeleteConmorbilityService) DeleteConmorbilitySvc(ctx context.Context, Id string) (deleteConmorbility.DeleteConmorbilityResponse, error) {
	d.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := d.repoDB.DeleteConmorbilityRepo(ctx, int64(IdConverter))
	if err != nil {
		d.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return deleteConmorbility.DeleteConmorbilityResponse{
			ResponseCode: http.StatusBadRequest,
			Messagge:     "failed",
		}, constants.ErrorDataError

		if resp == false {
			d.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return deleteConmorbility.DeleteConmorbilityResponse{
				ResponseCode: http.StatusBadRequest,
				Messagge:     "failed",
			}, constants.ErrorDataError
		}
	}
	return deleteConmorbility.DeleteConmorbilityResponse{
		ResponseCode: http.StatusOK,
		Messagge:     "Successful",
	}, nil
}
