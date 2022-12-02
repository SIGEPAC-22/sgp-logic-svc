package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/internal/updatePatientFile"
	"sgp-logic-svc/kit/constants"
	"strconv"
)

type UpdatePatientFileService struct {
	repoDB updatePatientFile.Repository
	logger kitlog.Logger
}

func NewUpdatePatientFileService(repoDB updatePatientFile.Repository, logger kitlog.Logger) *UpdatePatientFileService {
	return &UpdatePatientFileService{repoDB: repoDB, logger: logger}
}

func (u UpdatePatientFileService) UpdatePatientFileSvc(ctx context.Context, idPatient string, idPatientFile string, statePatient string, highDate string, lowDate string, comorbidity []string, symptom []string) (updatePatientFile.UpdatePatientFileResponse, error) {
	u.logger.Log("Starting Update Info Patient", constants.UUID, ctx.Value(constants.UUID))

	idPatientConvert, _ := strconv.Atoi(idPatient)
	idPatientFileConvert, _ := strconv.Atoi(idPatientFile)

	respSelectCbx, errSelect := u.repoDB.SelectPatientFileCBXRepo(ctx, idPatientConvert, idPatientFileConvert)
	if errSelect != nil {
		
	}
	//SELECT PARA IR A TRAER LA INFO DE LOS CBX

	//UPDATE EXCLUDED SYMPTOM COMORBIDITY

	//SELECT PARA TRAER LOS SYMPTOM RELACIONADOS AL ID ID

	//"GRIPE, TOS"
	//"GRIPE, TOS"

	//LOGICA MATCHpos
	//CREATE OR DELETE

	//SELECT PARA TRAER LOS COMORBIDITY RELACIONADOS AL ID ID

	//LOGICA MATCH
	//CREATE OR DELETE
	return updatePatientFile.UpdatePatientFileResponse{}, nil
}
