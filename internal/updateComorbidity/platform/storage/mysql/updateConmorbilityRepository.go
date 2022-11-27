package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	"sgp-logic-svc/kit/constants"
)

type UpdateComorbidityRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewUpdateComorbidityRepository(db *sql.DB, logger kitlog.Logger) *UpdateComorbidityRepository {
	return &UpdateComorbidityRepository{db: db, logger: logger}
}

func (u UpdateComorbidityRepository) UpdateComorbidityRepo(ctx context.Context, Id int64, NameComorbidity string, DescriptionComorbidity string) (bool, error) {
	sql, err := u.db.ExecContext(ctx, "UPDATE cby_comorbidity SET cby_name_comorbidity = ?, cby_comorbidity_description = ? WHERE cby_id_comorbidity = ?;", NameComorbidity, DescriptionComorbidity, Id)
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
