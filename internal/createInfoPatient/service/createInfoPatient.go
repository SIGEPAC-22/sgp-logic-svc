package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/createInfoPatient"
	"sgp-logic-svc/kit/constants"
)

type CreateInfoPatientSvc struct {
	repoDB createInfoPatient.Repository
	logger kitlog.Logger
}

func NewCreateInfoPatientSvc(repoDB createInfoPatient.Repository, logger kitlog.Logger) *CreateInfoPatientSvc {
	return &CreateInfoPatientSvc{repoDB: repoDB, logger: logger}
}

func (c CreateInfoPatientSvc) CreateInfoPatientSvc(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (createInfoPatient.CreateInfoPatientResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	resp, err := c.repoDB.CreateInfoPatientRepo(ctx, firstName, secondName, lastFirstName, lastSecondName, dateBirth, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, country, patientFile, patientSex)
	if err != nil {
		c.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createInfoPatient.CreateInfoPatientResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			c.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return createInfoPatient.CreateInfoPatientResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return createInfoPatient.CreateInfoPatientResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
