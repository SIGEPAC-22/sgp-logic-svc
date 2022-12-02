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

func (u *UpdatePatientFileRepository) SelectPatientFileRepo(ctx context.Context, id int) (updatePatientFile.SelectPatientFileResponse, error) {
	rows := u.db.QueryRowContext(ctx, "SELECT pat_id_document_type, pat_id_department, pfl_pregnant FROM sgp_info_svc.pat_patient\ninner join pfl_patient_file AS pfl\non pfl_id_patient = pat_id_patient\nWHERE pat_id_patient = ?;", id)
	u.logger.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))

	var respDB sqlGetUpdatePatientFile

	if err := rows.Scan(&respDB.DocumentType, &respDB.Department, &respDB.Pregnant); err != nil {
		u.logger.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return updatePatientFile.SelectPatientFileResponse{}, errors.New("Data not found")
	}
	resp := updatePatientFile.SelectPatientFileResponse{
		DocumentType: respDB.DocumentType,
		Department:   respDB.Department,
		Pregnant:     respDB.Pregnant,
	}

	return resp, nil
}
func (u UpdatePatientFileRepository) UpdatePatientFileRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "UPDATE pat_patient SET pat_first_name = ?, pat_second_name = ?, pat_first_last_name = ?, pat_second_last_name = ?, pat_id_document_type = ?, pat_document_number = ?,pat_cellphone_number = ?,pat_phone_number = ?,pat_reponsible_family = ?, pat_responsible_family_phone_number = ?,pat_id_department = ? WHERE pat_id_patient = ?;", firstName, secondName, lastFirstName, lastSecondName, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, Id)
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
