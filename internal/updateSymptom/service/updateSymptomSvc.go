package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/updateSymptom"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type UpdateSymptomService struct {
	repoDB updateSymptom.Repository
	logger kitlog.Logger
}

func NewUpdateSymptomService(repoDB updateSymptom.Repository, logger kitlog.Logger) *UpdateSymptomService {
	return &UpdateSymptomService{repoDB: repoDB, logger: logger}
}

func (u UpdateSymptomService) UpdateSymptomSvc(ctx context.Context, Id string, NameSymptom string, DescriptionSymptom string) (updateSymptom.UpdateSymptomResponse, error) {
	u.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	IdConverter, _ := strconv.Atoi(Id)

	resp, err := u.repoDB.UpdateSymptomRepo(ctx, int64(IdConverter), NameSymptom, DescriptionSymptom)
	if err != nil {
		u.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updateSymptom.UpdateSymptomResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			u.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return updateSymptom.UpdateSymptomResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return updateSymptom.UpdateSymptomResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
