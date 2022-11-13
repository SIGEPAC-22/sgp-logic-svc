package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/deleteComorbidity"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type DeleteComorbidityService struct {
	repoDB deleteComorbidity.Repository
	logger kitlog.Logger
}

func NewDeleteComorbidityService(repoDB deleteComorbidity.Repository, logger kitlog.Logger) *DeleteComorbidityService {
	return &DeleteComorbidityService{repoDB: repoDB, logger: logger}
}

func (d DeleteComorbidityService) DeleteComorbiditySvc(ctx context.Context, Id string) (deleteComorbidity.DeleteComorbidityResponse, error) {
	d.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := d.repoDB.DeleteComorbidityRepo(ctx, int64(IdConverter))
	if err != nil {
		d.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return deleteComorbidity.DeleteComorbidityResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			d.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return deleteComorbidity.DeleteComorbidityResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return deleteComorbidity.DeleteComorbidityResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
