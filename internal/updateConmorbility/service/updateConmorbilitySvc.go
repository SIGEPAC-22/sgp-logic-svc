package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/updateConmorbility"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type UpdateConmorbilityService struct {
	repoDB updateConmorbility.Repository
	logger kitlog.Logger
}

func NewUpdateConmorbilityService(repoDB updateConmorbility.Repository, logger kitlog.Logger) *UpdateConmorbilityService {
	return &UpdateConmorbilityService{repoDB: repoDB, logger: logger}
}

func (u UpdateConmorbilityService) UpdateConmorbilitySvc(ctx context.Context, Id string, NameConmorbility string, DescriptionConmorbility string) (updateConmorbility.UpdateUpdateConmorbilityResponse, error) {
	u.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := u.repoDB.UpdateConmorbilityRepo(ctx, int64(IdConverter), NameConmorbility, DescriptionConmorbility)
	if err != nil {
		u.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updateConmorbility.UpdateUpdateConmorbilityResponse{
			ResponseCode: http.StatusBadRequest,
			Messagge:     "failed",
		}, constants.ErrorDataError

		if resp == false {
			u.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return updateConmorbility.UpdateUpdateConmorbilityResponse{
				ResponseCode: http.StatusBadRequest,
				Messagge:     "failed",
			}, constants.ErrorDataError
		}
	}
	return updateConmorbility.UpdateUpdateConmorbilityResponse{
		ResponseCode: http.StatusOK,
		Messagge:     "Successful",
	}, nil
}
