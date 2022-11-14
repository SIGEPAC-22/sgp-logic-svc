package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/kit/constants"
)

type CreateInfoPatientRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewCreateInfoPatientRepository(db *sql.DB, logger kitlog.Logger) *CreateInfoPatientRepository {
	return &CreateInfoPatientRepository{db: db, logger: logger}
}

func (c CreateInfoPatientRepository) CreateInfoPatientRepo(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, country int, patientFile int, patientSex int) (bool, error) {

	sql, err := c.db.ExecContext(ctx, "INSERT INTO pat_patient (pat_first_name, pat_second_name, pat_first_last_name, pat_second_last_name, pat_date_birth, pat_id_document_type,\npat_document_number, pat_cellphone_number, pat_phone_number, pat_reponsible_family, pat_responsible_family_phone_number,\npat_id_department, pat_id_department_id_country, pat_id_patient_file, pat_id_patient_sex) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", firstName, secondName, lastFirstName, lastSecondName, dateBirth, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, country, patientFile, patientSex)
	c.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		c.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			c.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
