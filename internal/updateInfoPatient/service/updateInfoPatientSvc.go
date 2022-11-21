package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
	"sgp-logic-svc/internal/updateInfoPatient"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type UpdateInfoPatientService struct {
	repoDB updateInfoPatient.Repository
	logger kitlog.Logger
}

func NewUpdateInfoPatientService(repoDB updateInfoPatient.Repository, logger kitlog.Logger) *UpdateInfoPatientService {
	return &UpdateInfoPatientService{repoDB: repoDB, logger: logger}
}

//func (u UpdateInfoPatientService) UpdateInfoPatientSvc(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (updateInfoPatient.UpdateInfoPatientResponse, error) {
func (u UpdateInfoPatientService) UpdateInfoPatientSvc(ctx context.Context, Id string, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType string, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department string, country string) (updateInfoPatient.UpdateInfoPatientResponse, error) {
	u.logger.Log("Starting Update Info Patient", constants.UUID, ctx.Value(constants.UUID))

	idConvert, _ := strconv.Atoi(Id)
	idDocumentType, _ := strconv.Atoi(documentType)
	idDepartment, _ := strconv.Atoi(department)
	idCountry, _ := strconv.Atoi(country)

	//resp, err := u.repoDB.UpdateInfoPatientRepo(ctx, Id, firstName, secondName, lastFirstName, lastSecondName, dateBirth, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, country, patientFile, patientSex)
	resp, err := u.repoDB.UpdateInfoPatientRepo(ctx, idConvert, firstName, secondName, lastFirstName, lastSecondName, idDocumentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, idDepartment, idCountry)
	if err != nil {
		u.logger.Log("Error trying to push repository subscription", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updateInfoPatient.UpdateInfoPatientResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError

		if resp == false {
			u.logger.Log("No affected rows", constants.UUID, ctx.Value(constants.UUID))
			return updateInfoPatient.UpdateInfoPatientResponse{
				ResponseCode: http.StatusBadRequest,
				Message:      "failed",
			}, constants.ErrorDataError
		}
	}
	return updateInfoPatient.UpdateInfoPatientResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
