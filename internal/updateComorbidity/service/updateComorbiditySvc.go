package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/updateComorbidity"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type UpdateComorbidityService struct {
	repoDB updateComorbidity.Repository
	logger kitlog.Logger
}

func NewUpdateComorbidityService(repoDB updateComorbidity.Repository, logger kitlog.Logger) *UpdateComorbidityService {
	return &UpdateComorbidityService{repoDB: repoDB, logger: logger}
}

func (u UpdateComorbidityService) UpdateComorbiditySvc(ctx context.Context, Id string, NameComorbidity string, DescriptionComorbidity string) (updateComorbidity.UpdateComorbidityResponse, error) {
	u.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := u.repoDB.UpdateComorbidityRepo(ctx, int64(IdConverter), NameComorbidity, DescriptionComorbidity)
	if err != nil {
		u.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updateComorbidity.UpdateComorbidityResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			u.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return updateComorbidity.UpdateComorbidityResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return updateComorbidity.UpdateComorbidityResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
