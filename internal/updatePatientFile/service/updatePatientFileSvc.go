package service

import (
	"context"
	kitlog "github.com/go-kit/log"
	"net/http"
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

func (u UpdatePatientFileService) UpdatePatientFileSvc(ctx context.Context, idPatient string, idPatientFile string, statePatient updatePatientFile.StatePatient, highDate string, lowDate string, comorbidity updatePatientFile.Comorbidity, symptom updatePatientFile.Symptom) (updatePatientFile.UpdatePatientFileResponse, error) {
	u.logger.Log("Starting Update Info Patient", constants.UUID, ctx.Value(constants.UUID))

	var statePatientID int
	idPatientConvert, _ := strconv.Atoi(idPatient)
	idPatientFileConvert, _ := strconv.Atoi(idPatientFile)
	statePatientID, _ = strconv.Atoi(statePatient.Value)

	respSelectCbx, errSelect := u.repoDB.SelectPatientFileCBXRepo(ctx, idPatientConvert, idPatientFileConvert)
	if errSelect != nil {
		u.logger.Log("Error failed repo select Combobox", constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.UpdatePatientFileResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError
	}
	if statePatient.Value == "" {
		statePatientID = respSelectCbx.StatePatient
	}
	if highDate == "" {
		highDate = respSelectCbx.HighDate
	}
	if lowDate == "" {
		lowDate = respSelectCbx.LowDate
	}
	_, errUpdate := u.repoDB.UpdatePatientFileRepo(ctx, idPatientConvert, idPatientFileConvert, statePatientID, highDate, lowDate)
	if errUpdate != nil {
		u.logger.Log("Error failed repo select Combobox", constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.UpdatePatientFileResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//comorbidity

	respSelectComorbidity, errSelectComorbidity := u.repoDB.SelectPatientHasComorbidity(ctx, idPatientFileConvert)
	if errSelectComorbidity != nil {
		u.logger.Log("Error failed repo select patient has symptom", constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.UpdatePatientFileResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError
	}

	// INICIO TRANSFORMACION ARREGLO DE STRING A INT (COMORBIDITY)
	var arrayTransformReqIntComorbidity []int
	for _, arrayRequest := range comorbidity {
		dataInt, err := strconv.Atoi(arrayRequest.Value)
		if err != nil {
			panic(err)
		}
		arrayTransformReqIntComorbidity = append(arrayTransformReqIntComorbidity, dataInt)
	}
	// FIN TRANSFORMACION ARREGLO DE STRING A INT (COMORBIDITY)

	var arrayCreateComorbidity, arrayDeleteComorbidity []int
	if arrayTransformReqIntComorbidity != nil {
		if len(arrayTransformReqIntComorbidity) > len(respSelectComorbidity) {
			for i := 0; i < len(arrayTransformReqIntComorbidity); i++ {
				var num int
				for j := 0; j < len(respSelectComorbidity); j++ {
					if arrayTransformReqIntComorbidity[i] == respSelectComorbidity[j] {
						i += 1
						continue
					}
				}
				num = arrayTransformReqIntComorbidity[i]
				arrayCreateComorbidity = append(arrayCreateComorbidity, num)
			}
			//se valida que el request sea mayor a el de la db para que sea eliminacion
		} else if len(arrayTransformReqIntComorbidity) < len(respSelectComorbidity) {
			for i := 0; i < len(respSelectComorbidity); i++ {
				var num int
				for j := 0; j < len(arrayTransformReqIntComorbidity); j++ {
					if respSelectComorbidity[i] == arrayTransformReqIntComorbidity[j] {
						i += 1
						continue
					}
				}
				num = respSelectComorbidity[i]
				arrayDeleteComorbidity = append(arrayDeleteComorbidity, num)
			}
		}

		if len(arrayCreateComorbidity) > 0 {
			for _, arrayCreateList := range arrayCreateComorbidity {
				_, errResp := u.repoDB.CreatePatientComorbidity(ctx, arrayCreateList, idPatientFileConvert)
				if errResp != nil {

					continue
				}
			}
		}

		if len(arrayDeleteComorbidity) > 0 {
			for _, arrayDeleteList := range arrayDeleteComorbidity {
				_, errResp := u.repoDB.DeletePatientComorbidity(ctx, arrayDeleteList, idPatientFileConvert)
				if errResp != nil {

				}
			}
		}
	}
	//comorbidity
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//symptom

	respSelectSymptom, errSelectSymptom := u.repoDB.SelectPatientHasSymptom(ctx, idPatientFileConvert)
	if errSelectSymptom != nil {
		u.logger.Log("Error failed repo select patient has symptom", constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.UpdatePatientFileResponse{
			ResponseCode: http.StatusBadRequest,
			Message:      "failed",
		}, constants.ErrorDataError
	}

	// INICIO TRANSFORMACION ARREGLO DE STRING A INT (SYMPTOM)
	var arrayTransformReqIntSymptom []int
	for _, arrayRequest := range symptom {
		dataInt, err := strconv.Atoi(arrayRequest.Value)
		if err != nil {
			panic(err)
		}
		arrayTransformReqIntSymptom = append(arrayTransformReqIntSymptom, dataInt)
	}
	// FIN TRANSFORMACION ARREGLO DE STRING A INT (SYMPTOM)

	var arrayCreateSymptom, arrayDeleteSymptom []int
	if arrayTransformReqIntSymptom != nil {
		if len(arrayTransformReqIntSymptom) > len(respSelectSymptom) {
			for i := 0; i < len(arrayTransformReqIntSymptom); i++ {
				var num int
				for j := 0; j < len(respSelectSymptom); j++ {
					if arrayTransformReqIntSymptom[i] == respSelectSymptom[j] {
						i += 1
						continue
					}
				}
				num = arrayTransformReqIntSymptom[i]
				arrayCreateSymptom = append(arrayCreateSymptom, num)
			}
			//se valida que el request sea mayor a el de la db para que sea eliminacion
		} else if len(arrayTransformReqIntSymptom) < len(respSelectSymptom) {
			for i := 0; i < len(respSelectSymptom); i++ {
				var num int
				for j := 0; j < len(arrayTransformReqIntSymptom); j++ {
					if respSelectSymptom[i] == arrayTransformReqIntSymptom[j] {
						i += 1
						continue
					}
				}
				num = respSelectSymptom[i]
				arrayDeleteSymptom = append(arrayDeleteSymptom, num)
			}
		}

		if len(arrayCreateSymptom) > 0 {
			for _, arrayCreateList := range arrayCreateSymptom {
				_, errResp := u.repoDB.CreatePatientSymptom(ctx, arrayCreateList, idPatientFileConvert)
				if errResp != nil {

					continue
				}
			}
		}

		if len(arrayDeleteSymptom) > 0 {
			for _, arrayDeleteList := range arrayDeleteSymptom {
				_, errResp := u.repoDB.DeletePatientSymptom(ctx, arrayDeleteList, idPatientFileConvert)
				if errResp != nil {

				}
			}
		}
	}
	//symptom
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	return updatePatientFile.UpdatePatientFileResponse{
		ResponseCode: http.StatusOK,
		Message:      "Successful",
	}, nil
}
