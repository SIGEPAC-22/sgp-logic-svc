package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/createInfoPatient"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type CreateInfoPatientSvc struct {
	repoDB createInfoPatient.Repository
	logger kitlog.Logger
}

func NewCreateInfoPatientSvc(repoDB createInfoPatient.Repository, logger kitlog.Logger) *CreateInfoPatientSvc {
	return &CreateInfoPatientSvc{repoDB: repoDB, logger: logger}
}

func (c CreateInfoPatientSvc) CreateInfoPatientSvc(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType string, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department string, patientSex string, pregnant string, foreign string) (createInfoPatient.CreateInfoPatientResponse, error) {
	c.logger.Log("Starting subscription", constants.UUID, ctx.Value(constants.UUID))

	convertDocumentType, _ := strconv.Atoi(documentType)
	convertDepartment, _ := strconv.Atoi(department)
	convertPatientSex, _ := strconv.Atoi(patientSex)
	foreignInteger, _ := strconv.Atoi(foreign)
	//pregnantInteger, _ := strconv.Atoi(foreign)

	var pregnantBoolean bool
	//var foreignInteger int

	if pregnant == "1" {
		pregnantBoolean = true
	} else if pregnant == "0" {
		pregnantBoolean = false
	}

	/*if foreign == "si" {
		foreignInteger = 2
	} else if foreign == "" {
		foreignInteger = 1
	}*/

	resp, err := c.repoDB.CreateInfoPatientRepo(ctx, firstName, secondName, lastFirstName, lastSecondName, dateBirth, convertDocumentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, convertDepartment, convertPatientSex, foreignInteger)
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

	respInfoPatient, errRespInfoPatient := c.repoDB.SelectInfoPatient(ctx, firstName, secondName, lastFirstName, documentNumber)
	if errRespInfoPatient != nil {
		c.logger.Log("Error obtaining patient information", "DocumentNumber", documentNumber, "Error", errRespInfoPatient, constants.UUID, ctx.Value(constants.UUID))
		return createInfoPatient.CreateInfoPatientResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, errRespInfoPatient
	}

	_, errPatientFile := c.repoDB.CreatePatientFileRepo(ctx, respInfoPatient, pregnantBoolean)
	if errPatientFile != nil {
		c.logger.Log("Error trying to push repository subscription", "error", errPatientFile.Error(), constants.UUID, ctx.Value(constants.UUID))
		return createInfoPatient.CreateInfoPatientResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, errPatientFile
	}

	return createInfoPatient.CreateInfoPatientResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
