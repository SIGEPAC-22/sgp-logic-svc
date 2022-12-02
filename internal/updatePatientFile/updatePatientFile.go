package updatePatientFile

import (
	"context"
)

type Repository interface {
	SelectPatientFileCBXRepo(ctx context.Context, id int) (SelectPatientFileCbxResponse, error)
	UpdatePatientFileRepo(ctx context.Context, idPatient string, idPatientFile string, statePatient string, highDate string, lowDate string) (bool, error)
	SelectPatientHasSymptom(ctx context.Context, idSymptom int, idPatientFile int) (SelectPatientSymptom, error)
	CreatePatientSymptom(ctx context.Context, idSymptom int, idPatientFile int)
	DeletePatientSymptom(ctx context.Context, idPatientFile int) (bool, error)
	SelectPatientHasComorbidity(ctx context.Context, idComorbidity int, idPatientFile int) (SelectPatientComorbidity, error)
	CreatePatientComorbidity(ctx context.Context, idComorbidity int, idPatientFile int)
	DeletePatientComorbidity(ctx context.Context, idPatientFile int) (bool, error)
}

type Service interface {
	UpdatePatientFileSvc(ctx context.Context, idPatient string, idPatientFile string, statePatient string, highDate string, lowDate string, comorbidity []string, symptom []string) (UpdatePatientFileResponse, error)
}

type UpdatePatientFileRequest struct {
	IdPatient     int      `json:"idPatient"`
	IdPatientFile int      `json:"idPatientFile"`
	StatePatient  string   `json:"statePatient"`
	HighDate      string   `json:"highDate"`
	LowDate       string   `json:"lowDate"`
	Comorbidity   []string `json:"comorbidity"`
	Symptom       []string `json:"symptom"`
}

type SelectPatientFileCbxResponse struct {
	StatePatient string `json:"statePatient"`
	HighDate     string `json:"highDate"`
	LowDate      string `json:"lowDate"`
}

type UpdatePatientFileResponse struct {
	ResponseCode int    `json:"responseCode"`
	Message      string `json:"message"`
}

type SelectPatientSymptom struct {
	SymptomPatientList []string `json:"symptomPatientList"`
}

type SelectPatientComorbidity struct {
	ComorbidityPatientList []string `json:"comorbidityPatientList"`
}
