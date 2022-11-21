package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/kit/constants"
)

type UpdateInfoPatientRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdateInfoPatientRepository(db *sql.DB, logger kitlog.Logger) *UpdateInfoPatientRepository {
	return &UpdateInfoPatientRepository{db: db, logger: logger}
}

func (u UpdateInfoPatientRepository) UpdateInfoPatientRepo(ctx context.Context, Id int, firstName string, secondName string, lastFirstName string, lastSecondName string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int) (bool, error) {
	//sql, err := u.db.ExecContext(ctx, "UPDATE pat_patient SET pat_first_name = ?, pat_second_name = ?, pat_first_last_name = ?, pat_second_last_name = ?, pat_date_birth = ?, pat_id_document_type = ?, pat_document_number = ?,pat_cellphone_number = ?,pat_phone_number = ?,pat_reponsible_family = ?,pat_responsible_family_phone_number = ?,pat_id_department = ?,pat_id_department_id_country = ?,pat_id_patient_file = ?,pat_id_patient_sex = ? WHERE pat_id_patient = ?;", firstName, secondName, lastFirstName, lastSecondName, dateBirth, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, country, patientFile, patientSex, Id)
	sql, err := u.db.ExecContext(ctx, "UPDATE pat_patient SET pat_first_name = ?, pat_second_name = ?, pat_first_last_name = ?, pat_second_last_name = ?, pat_id_document_type = ?, pat_document_number = ?,pat_cellphone_number = ?,pat_phone_number = ?,pat_reponsible_family = ?, pat_responsible_family_phone_number = ?,pat_id_department = ?,pat_id_country = ? WHERE pat_id_patient = ?;", firstName, secondName, lastFirstName, lastSecondName, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, country, Id)
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
