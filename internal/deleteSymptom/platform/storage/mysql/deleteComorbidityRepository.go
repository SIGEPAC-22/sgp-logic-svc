package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-logic-svc/kit/constants"
)

type DeleteSymptomRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewDeleteSymptomRepository(db *sql.DB, logger kitlog.Logger) *DeleteSymptomRepository {
	return &DeleteSymptomRepository{db: db, logger: logger}
}

func (d DeleteSymptomRepository) DeleteSymptomRepo(ctx context.Context, Id int64) (bool, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	StatusInactive := config.GetInt("appProperties.deleteComorbidityStatusInactive")

	sql, err := d.db.ExecContext(ctx, "UPDATE stm_symptom SET stm_id_state_data = ? WHERE stm_id_sympton = ?;", StatusInactive, Id)
	d.logger.Log("query about to exec", "query", sql, constants.UUID, ctx.Value(constants.UUID))
	if err != nil {
		d.logger.Log("Error when trying to insert information", "error", err.Error(), constants.UUID, ctx.Value(constants.UUID))
		return false, err

		rows, _ := sql.RowsAffected()
		if rows != 1 {
			d.logger.Log("zero rows affected", constants.UUID, ctx.Value(constants.UUID))
			return false, err
		}
	}
	return true, nil
}
