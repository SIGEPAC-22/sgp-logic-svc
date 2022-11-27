package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/log"
	goconfig "github.com/iglin/go-config"
	"sgp-logic-svc/kit/constants"
)

type CreateSymptomRepository struct {
	db     *sql.DB
	logger kitlog.Logger
}

func NewCreateSymptomRepository(db *sql.DB, logger kitlog.Logger) *CreateSymptomRepository {
	return &CreateSymptomRepository{db: db, logger: logger}
}

func (c CreateSymptomRepository) CreateSymptomRepo(ctx context.Context, NameSymptom string, DescriptionSymptom string) (bool, error) {

	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	statusInitial := config.GetInt("appProperties.createComorbidityStatusInitial")
	sql, err := c.db.ExecContext(ctx, "INSERT INTO stm_symptom (stm_name_symptons,stm_sympton_description,stm_id_state_data) VALUES (?,?,?);", NameSymptom, DescriptionSymptom, statusInitial)
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
