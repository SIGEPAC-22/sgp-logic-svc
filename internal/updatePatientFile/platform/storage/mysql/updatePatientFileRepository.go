package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/internal/updatePatientFile"
	"sgp-logic-svc/kit/constants"
)

type UpdatePatientFileRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdatePatientFileRepository(db *sql.DB, logger kitlog.Logger) *UpdatePatientFileRepository {
	return &UpdatePatientFileRepository{db: db, logger: logger}
}

func (u UpdatePatientFileRepository) SelectPatientFileCBXRepo(ctx context.Context, idPatient int, idPatientFile int) (updatePatientFile.SelectPatientFileCbxResponse, error) {
	rows := u.db.QueryRowContext(ctx, "SELECT pfl_high_date,pfl_low_date,pfl_id_state_patient FROM pfl_patient_file where pfl_id_patient_file = ? and pfl_id_patient=?;", idPatient, idPatientFile)
	u.logger.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))

	var respBD sqlGetUpdatePatientFile
	if err := rows.Scan(&respBD.HighDate, &respBD.LowDate, &respBD.StatePatient); err != nil {
		u.logger.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.SelectPatientFileCbxResponse{}, errors.New("Data not found")
	}
	resp := updatePatientFile.SelectPatientFileCbxResponse{
		StatePatient: respBD.StatePatient,
		HighDate:     respBD.HighDate,
		LowDate:      respBD.LowDate,
	}
	return resp, nil
}

func (u UpdatePatientFileRepository) UpdatePatientFileRepo(ctx context.Context, idPatient int, idPatientFile int, statePatient string, highDate string, lowDate string) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "UPDATE pfl_patient_file SET pfl_id_state_patient = ?, pfl_high_date = ?, pfl_low_date = ? WHERE pfl_id_patient_file = ? AND pfl_id_patient = ?;", statePatient, highDate, lowDate, idPatientFile, idPatient)

	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to update information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err
		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
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
