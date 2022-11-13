package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sgp-logic-svc/internal/createComorbidity/platform/handler"
	"sgp-logic-svc/internal/createComorbidity/platform/storage/mysql"
	"sgp-logic-svc/internal/createComorbidity/service"
	handler3 "sgp-logic-svc/internal/deleteComorbidity/platform/handler"
	mysql3 "sgp-logic-svc/internal/deleteComorbidity/platform/storage/mysql"
	service3 "sgp-logic-svc/internal/deleteComorbidity/service"
	handler2 "sgp-logic-svc/internal/updateConmorbility/platform/handler"
	mysql2 "sgp-logic-svc/internal/updateConmorbility/platform/storage/mysql"
	service2 "sgp-logic-svc/internal/updateConmorbility/service"
	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)
	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)
	////////////////////////CORS///////////////////////////////////////////

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	//////////////////////CREATE CONMORBILITY////////////////////////////////////////////////
	createComorbidityRepo := mysql.NewCreateComorbidityRepository(db, kitlogger)
	createComorbidityService := service.NewCreateComorbiditySvc(createComorbidityRepo, kitlogger)
	createComorbidityEndpoint := handler.MakeCreateComorbidityEndpoint(createComorbidityService)
	createComorbidityEndpoint = handler.CreateComorbidityTransportMiddleware(kitlogger)(createComorbidityEndpoint)
	createComorbidityHandler := handler.NewCreateComorbidityHandler(config.GetString("paths.createComorbidity"), createComorbidityEndpoint)
	//////////////////////CREATE CONMORBILITY////////////////////////////////////////////////

	//////////////////////UPDATE CONMORBILITY////////////////////////////////////////////////
	updateConmorbilityRepo := mysql2.NewUpdateConmorbilityRepository(db, kitlogger)
	updateConmorbilityService := service2.NewUpdateConmorbilityService(updateConmorbilityRepo, kitlogger)
	updateConmorbilityEndpoint := handler2.MakeUpdateUpdateConmorbilityEndpoint(updateConmorbilityService)
	updateConmorbilityEndpoint = handler2.UpdateConmorbilityTransportMiddleware(kitlogger)(updateConmorbilityEndpoint)
	updateConmorbilityHandler := handler2.NewUpdateConmorbilityHandler(config.GetString("paths.updateConmorbility"), updateConmorbilityEndpoint)
	//////////////////////UPDATE CONMORBILITY////////////////////////////////////////////////

	//////////////////////DELETE CONMORBILITY////////////////////////////////////////////////
	deleteConmorbilityRepo := mysql3.NewDeleteConmorbilityRepository(db, kitlogger)
	deleteConmorbilityService := service3.NewDeleteConmorbilityService(deleteConmorbilityRepo, kitlogger)
	deleteConmorbilityEndpoint := handler3.MakeDeleteConmorbilityEndpoint(deleteConmorbilityService)
	deleteConmorbilityEndpoint = handler3.DeleteConmorbilityTransportMiddleware(kitlogger)(deleteConmorbilityEndpoint)
	deleteConmorbilityHandler := handler3.NewDeleteConmorbilityHandler(config.GetString("paths.deleteComorbidity"), deleteConmorbilityEndpoint)
	//////////////////////DELETE CONMORBILITY////////////////////////////////////////////////

	mux.Handle(config.GetString("paths.createComorbidity"), createComorbidityHandler)
	mux.Handle(config.GetString("paths.updateComorbidity"), updateConmorbilityHandler)
	mux.Handle(config.GetString("paths.deleteComorbidity"), deleteConmorbilityHandler)
	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
