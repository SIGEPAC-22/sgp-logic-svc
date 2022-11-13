package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-logic-svc/kit/constants"
)

type DeleteConmorbilityRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewDeleteConmorbilityRepository(db *sql.DB, logger kitlog.Logger) *DeleteConmorbilityRepository {
	return &DeleteConmorbilityRepository{db: db, logger: logger}
}

func (d DeleteConmorbilityRepository) DeleteConmorbilityRepo(ctx context.Context, Id int64) (bool, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	StatusInactive := config.GetInt("appProperties.deleteConmorbilityStatusInactive")

	sql, err := d.db.ExecContext(ctx, "UPDATE cby_comorbidity SET cby_id_state_data = ? WHERE cby_id_comorbidity = ?;", StatusInactive, Id)
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
