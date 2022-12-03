package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-logic-svc/internal/updatePatientFile"
	"sgp-logic-svc/kit/constants"
	"time"
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
		HighDate:     transformerPointer(respBD.HighDate),
		LowDate:      transformerPointer(respBD.LowDate),
	}
	return resp, nil
}

func (u UpdatePatientFileRepository) UpdatePatientFileRepo(ctx context.Context, idPatient int, idPatientFile int, statePatient int, highDate string, lowDate string) (bool, error) {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	var status, statusHigh, statusLow bool

	status = config.GetBool("app-PropertiesUpdateBot.false")
	statusHigh = config.GetBool("app-PropertiesUpdateBot.false")
	statusLow = config.GetBool("app-PropertiesUpdateBot.false")

	if highDate != "0001-01-01" {
		status = config.GetBool("app-PropertiesUpdateBot.true")
		statusHigh = config.GetBool("app-PropertiesUpdateBot.true")
		statusLow = config.GetBool("app-PropertiesUpdateBot.false")
	} else if lowDate != "0001-01-01" {
		status = config.GetBool("app-PropertiesUpdateBot.true")
		statusHigh = config.GetBool("app-PropertiesUpdateBot.false")
		statusLow = config.GetBool("app-PropertiesUpdateBot.true")
	}

	sql, err := u.db.ExecContext(ctx, "UPDATE pfl_patient_file SET pfl_id_state_patient = ?, pfl_high = ?, pfl_high_date = ?,pfl_low = ?, pfl_low_date = ?, pft_status_bot = ? WHERE pfl_id_patient_file = ? AND pfl_id_patient = ?;", statePatient, statusHigh, highDate, statusLow, lowDate, status, idPatientFile, idPatient)

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

func (u UpdatePatientFileRepository) SelectPatientHasSymptom(ctx context.Context, idPatientFile int) ([]int, error) {

	var resp []int
	rows, errDB := u.db.QueryContext(ctx, "SELECT fhs_id_symptom FROM fhs_file_has_sympton WHERE fhs_id_patient_file = ?;", idPatientFile)
	if errDB != nil {
		u.logger.Log("Error while trying to get information for Symptom-PatientFile", constants.UUID, ctx.Value(constants.UUID))
		return nil, errDB
	}

	defer rows.Close()
	for rows.Next() {
		var respDB SqlGetSymptomPatient
		if err := rows.Scan(&respDB.IdSymptom); err != nil {
			u.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return nil, err
		}

		resp = append(resp, respDB.IdSymptom)

	}
	return resp, nil
}

func (u UpdatePatientFileRepository) CreatePatientSymptom(ctx context.Context, idSymptom int, idPatientFile int) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "INSERT INTO fhs_file_has_sympton (fhs_id_symptom,fhs_id_patient_file) VALUES(?,?);", idSymptom, idPatientFile)
	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}

func (u UpdatePatientFileRepository) DeletePatientSymptom(ctx context.Context, idSymptom, idPatientFile int) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "DELETE FROM fhs_file_has_sympton WHERE  fhs_id_symptom = ? AND fhs_id_patient_file = ?;", idSymptom, idPatientFile)
	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}

func (u UpdatePatientFileRepository) SelectPatientHasComorbidity(ctx context.Context, idPatientFile int) ([]int, error) {
	var resp []int
	rows, errDB := u.db.QueryContext(ctx, "SELECT fhc_id_conmorbilities FROM fhc_file_has_cormobility WHERE fhc_id_patient_file = ?;", idPatientFile)
	if errDB != nil {
		u.logger.Log("Error while trying to get information for Symptom-PatientFile", constants.UUID, ctx.Value(constants.UUID))
		return nil, errDB
	}

	defer rows.Close()
	for rows.Next() {
		var respDB SqlGetSymptomPatient
		if err := rows.Scan(&respDB.IdSymptom); err != nil {
			u.logger.Log("error while trying to scan response from DB", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
			return nil, err
		}

		resp = append(resp, respDB.IdSymptom)

	}
	return resp, nil
}

func (u UpdatePatientFileRepository) CreatePatientComorbidity(ctx context.Context, idComorbidity int, idPatientFile int) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "INSERT INTO fhc_file_has_cormobility (fhc_id_conmorbilities,fhc_id_patient_file) VALUES(?,?);", idComorbidity, idPatientFile)
	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}

func (u UpdatePatientFileRepository) DeletePatientComorbidity(ctx context.Context, idComorbidity, idPatientFile int) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "DELETE FROM fhc_file_has_cormobility WHERE  fhc_id_conmorbilities = ? AND fhc_id_patient_file = ?;", idComorbidity, idPatientFile)
	u.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		u.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			u.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
func transformerPointer(date *time.Time) string {
	if date != nil {
		var dateConverter string

		config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
		dateConverter = date.Format(config.GetString("app-PropertiesUpdateBot.date"))
		return dateConverter
	} else {
		return "0000-00-00"
	}
}
