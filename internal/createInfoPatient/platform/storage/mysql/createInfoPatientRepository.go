package mysql

import (
	"context"
	"database/sql"
	"errors"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/kit/constants"
	"time"
)

type CreateInfoPatientRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewCreateInfoPatientRepository(db *sql.DB, logger kitlog.Logger) *CreateInfoPatientRepository {
	return &CreateInfoPatientRepository{db: db, logger: logger}
}

func (c *CreateInfoPatientRepository) CreateInfoPatientRepo(ctx context.Context, firstName string, secondName string, lastFirstName string, lastSecondName string, dateBirth string, documentType int, documentNumber string, cellphoneNumber string, phoneNumber string, responsibleFamily string, responsibleFamilyPhoneNumber string, department int, patientSex int, foreign int) (bool, error) {

	sql, err := c.db.ExecContext(ctx, "INSERT INTO pat_patient\n(pat_first_name,\npat_second_name,\npat_first_last_name,\npat_second_last_name,\npat_date_birth,\npat_id_document_type,\npat_document_number,\npat_cellphone_number,\npat_phone_number,\npat_reponsible_family,\npat_responsible_family_phone_number,\npat_id_department, pat_id_country,\npat_id_patient_sex) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?);", firstName, secondName, lastFirstName, lastSecondName, dateBirth, documentType, documentNumber, cellphoneNumber, phoneNumber, responsibleFamily, responsibleFamilyPhoneNumber, department, foreign, patientSex)
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

func (c *CreateInfoPatientRepository) SelectInfoPatient(ctx context.Context, firstName string, secondName string, lastFirstName string, documentNumber string) (int, error) {

	rows := c.db.QueryRowContext(ctx, "SELECT pat_id_patient FROM pat_patient where pat_first_name = ? AND pat_second_name = ? AND pat_first_last_name = ? AND pat_document_number = ?;", firstName, secondName, lastFirstName, documentNumber)
	c.logger.Log("query about so exec select", "query", rows, constants.UUID, ctx.Value(constants.UUID))

	var respDB sqlSelectInfoPatient

	if err := rows.Scan(&respDB.Id); err != nil {
		c.logger.Log("Data not found", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return 0, errors.New("Data not found")
	}
	resp := respDB.Id

	return resp, nil
}

func (c *CreateInfoPatientRepository) CreatePatientFileRepo(ctx context.Context, id int, pregnat bool) (bool, error) {

	dateNow := time.Now()

	sql, err := c.db.ExecContext(ctx, "INSERT INTO pfl_patient_file(pfl_admission_date, pfl_pregnant, pfl_id_patient) VALUES(?,?,?);", dateNow, pregnat, id)
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
