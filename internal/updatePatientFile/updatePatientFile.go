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
	UpdatePatientFileSvc(ctx context.Context, idPatient string, idPatientFile string, statePatient StatePatient, highDate string, lowDate string, comorbidity Comorbidity, symptom Symptom) (UpdatePatientFileResponse, error)
}

type UpdatePatientFileRequest struct {
	IdPatient     int          `json:"idPatient"`
	IdPatientFile int          `json:"idPatientFile"`
	StatePatient  StatePatient `json:"statePatient"`
	HighDate      string       `json:"highDate"`
	LowDate       string       `json:"lowDate"`
	Comorbidity   Comorbidity  `json:"comorbidity"`
	Symptom       Symptom      `json:"symptom"`
}

type Comorbidity []struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Symptom []struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type StatePatient struct {
	Value string `json:"value"`
	Label string `json:"label"`
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
