package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/kit/constants"
)

type UpdateSymptomRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdateSymptomRepository(db *sql.DB, logger kitlog.Logger) *UpdateSymptomRepository {
	return &UpdateSymptomRepository{db: db, logger: logger}
}

func (u UpdateSymptomRepository) UpdateSymptomRepo(ctx context.Context, Id int64, NameSymptom string, DescriptionSymptom string) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "UPDATE stm_symptom SET stm_name_symptons = ?, stm_sympton_description = ? WHERE stm_id_sympton = ?;", NameSymptom, DescriptionSymptom, Id)
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
