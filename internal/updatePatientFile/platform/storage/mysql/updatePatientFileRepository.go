package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/internal/updatePatientFile"
)

type UpdatePatientFileRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdatePatientFileRepository(db *sql.DB, logger kitlog.Logger) *UpdatePatientFileRepository {
	return &UpdatePatientFileRepository{db: db, logger: logger}
}

func (u UpdatePatientFileRepository) SelectPatientFileCBXRepo(ctx context.Context, id int) (updatePatientFile.SelectPatientFileCbxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) UpdatePatientFileRepo(ctx context.Context, idPatient string, idPatientFile string, statePatient string, highDate string, lowDate string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) SelectPatientHasSymptom(ctx context.Context, idSymptom int, idPatientFile int) (updatePatientFile.SelectPatientSymptom, error) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) CreatePatientSymptom(ctx context.Context, idSymptom int, idPatientFile int) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) DeletePatientSymptom(ctx context.Context, idPatientFile int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) SelectPatientHasComorbidity(ctx context.Context, idComorbidity int, idPatientFile int) (updatePatientFile.SelectPatientComorbidity, error) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) CreatePatientComorbidity(ctx context.Context, idComorbidity int, idPatientFile int) {
	//TODO implement me
	panic("implement me")
}

func (u UpdatePatientFileRepository) DeletePatientComorbidity(ctx context.Context, idPatientFile int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
