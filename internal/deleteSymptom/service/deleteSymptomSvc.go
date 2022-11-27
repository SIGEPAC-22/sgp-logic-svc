package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/deleteSymptom"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type DeleteSymptomService struct {
	repoDB deleteSymptom.Repository
	logger kitlog.Logger
}

func NewDeleteSymptomService(repoDB deleteSymptom.Repository, logger kitlog.Logger) *DeleteSymptomService {
	return &DeleteSymptomService{repoDB: repoDB, logger: logger}
}

func (d DeleteSymptomService) DeleteSymptomSvc(ctx context.Context, Id string) (deleteSymptom.DeleteSymptomResponse, error) {
	d.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := d.repoDB.DeleteSymptomRepo(ctx, int64(IdConverter))
	if err != nil {
		d.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return deleteSymptom.DeleteSymptomResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			d.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return deleteSymptom.DeleteSymptomResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return deleteSymptom.DeleteSymptomResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
