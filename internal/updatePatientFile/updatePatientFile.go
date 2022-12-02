package updatePatientFile

import (
	"context"
)

type Repository interface {
	SelectPatientFileCBXRepo(ctx context.Context, idPatient int, idPatientFile int) (SelectPatientFileCbxResponse, error)
	UpdatePatientFileRepo(ctx context.Context, idPatient int, idPatientFile int, statePatient int, highDate string, lowDate string) (bool, error)
	SelectPatientHasSymptom(ctx context.Context, idPatientFile int) ([]int, error)
	CreatePatientSymptom(ctx context.Context, idSymptom int, idPatientFile int) (bool, error)
	DeletePatientSymptom(ctx context.Context, idSymptom, idPatientFile int) (bool, error)
	SelectPatientHasComorbidity(ctx context.Context, idPatientFile int) ([]int, error)
	CreatePatientComorbidity(ctx context.Context, idComorbidity int, idPatientFile int) (bool, error)
	DeletePatientComorbidity(ctx context.Context, idComorbidity, idPatientFile int) (bool, error)
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
	StatePatient int    `json:"statePatient"`
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
