package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-logic-svc/kit/constants"
)

type CreateConmorbilityRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewCreateConmorbilityRepository(db *sql.DB, logger kitlog.Logger) *CreateConmorbilityRepository {
	return &CreateConmorbilityRepository{db: db, logger: logger}
}

func (c CreateConmorbilityRepository) CreateConmorbilityRepo(ctx context.Context, NameConmorbility string, DescriptionConmorbility string) (bool, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	statusInitial := config.GetInt("appProperties.createConmorbilityStatusInitial")
	sql, err := c.db.ExecContext(ctx, "INSERT INTO cby_comorbidity (cby_name_comorbidity,cby_comorbidity_description,cby_id_state_data)VALUES(?,?,?);", NameConmorbility, DescriptionConmorbility, statusInitial)
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
