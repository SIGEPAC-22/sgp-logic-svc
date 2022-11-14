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
	handler7 "sgp-logic-svc/internal/createInfoPatient/platform/handler"
	mysql7 "sgp-logic-svc/internal/createInfoPatient/platform/storage/mysql"
	service7 "sgp-logic-svc/internal/createInfoPatient/service"
	handler4 "sgp-logic-svc/internal/createSymptom/platform/handler"
	mysql4 "sgp-logic-svc/internal/createSymptom/platform/storage/mysql"
	service4 "sgp-logic-svc/internal/createSymptom/service"
	handler3 "sgp-logic-svc/internal/deleteComorbidity/platform/handler"
	mysql3 "sgp-logic-svc/internal/deleteComorbidity/platform/storage/mysql"
	service3 "sgp-logic-svc/internal/deleteComorbidity/service"
	handler5 "sgp-logic-svc/internal/deleteSymptom/platform/handler"
	mysql5 "sgp-logic-svc/internal/deleteSymptom/platform/storage/mysql"
	service5 "sgp-logic-svc/internal/deleteSymptom/service"
	handler2 "sgp-logic-svc/internal/updateComorbidity/platform/handler"
	mysql2 "sgp-logic-svc/internal/updateComorbidity/platform/storage/mysql"
	service2 "sgp-logic-svc/internal/updateComorbidity/service"
	handler6 "sgp-logic-svc/internal/updateSymptom/platform/handler"
	mysql6 "sgp-logic-svc/internal/updateSymptom/platform/storage/mysql"
	service6 "sgp-logic-svc/internal/updateSymptom/service"
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

	//////////////////////CREATE SYMPTOM////////////////////////////////////////////////
	createSymptomRepo := mysql4.NewCreateSymptomRepository(db, kitlogger)
	createSymptomService := service4.NewCreateSymptomSvc(createSymptomRepo, kitlogger)
	createSymptomEndpoint := handler4.MakeCreateSymptomEndpoint(createSymptomService)
	createSymptomEndpoint = handler4.CreateSymptomTransportMiddleware(kitlogger)(createSymptomEndpoint)
	createSymptomHandler := handler4.NewCreateSymptomHandler(config.GetString("paths.createSymptom"), createSymptomEndpoint)
	//////////////////////CREATE SYMPTOM////////////////////////////////////////////////

	//////////////////////CREATE INFO PATIENT////////////////////////////////////////////////
	createInfoPatientRepo := mysql7.NewCreateInfoPatientRepository(db, kitlogger)
	createInfoPatientService := service7.NewCreateInfoPatientSvc(createInfoPatientRepo, kitlogger)
	createInfoPatientEndpoint := handler7.MakeCreateInfoPatientEndpoint(createInfoPatientService)
	createInfoPatientEndpoint = handler7.CreateInfoPatientTransportMiddleware(kitlogger)(createInfoPatientEndpoint)
	createInfoPatientHandler := handler7.NewCreateInfoPatientHandler(config.GetString("paths.createInfoPatient"), createInfoPatientEndpoint)
	//////////////////////CREATE INFO PATIENT////////////////////////////////////////////////

	//////////////////////UPDATE CONMORBILITY////////////////////////////////////////////////
	updateComorbidityRepo := mysql2.NewUpdateComorbidityRepository(db, kitlogger)
	updateComorbidityService := service2.NewUpdateComorbidityService(updateComorbidityRepo, kitlogger)
	updateComorbidityEndpoint := handler2.MakeUpdateComorbidityEndpoint(updateComorbidityService)
	updateComorbidityEndpoint = handler2.UpdateComorbidityTransportMiddleware(kitlogger)(updateComorbidityEndpoint)
	updateComorbidityHandler := handler2.NewUpdateComorbidityHandler(config.GetString("paths.updateComorbidity"), updateComorbidityEndpoint)
	//////////////////////UPDATE CONMORBILITY////////////////////////////////////////////////

	//////////////////////UPDATE SYMPTOM////////////////////////////////////////////////
	updateSymptomRepo := mysql6.NewUpdateSymptomRepository(db, kitlogger)
	updateSymptomService := service6.NewUpdateSymptomService(updateSymptomRepo, kitlogger)
	updateSymptomEndpoint := handler6.MakeUpdateSymptomEndpoint(updateSymptomService)
	updateSymptomEndpoint = handler6.UpdateSymptomTransportMiddleware(kitlogger)(updateSymptomEndpoint)
	updateSymptomHandler := handler6.NewUpdateSymptomHandler(config.GetString("paths.updateSymptom"), updateSymptomEndpoint)
	//////////////////////UPDATE SYMPTOM////////////////////////////////////////////////

	//////////////////////DELETE CONMORBILITY////////////////////////////////////////////////
	deleteComorbidityRepo := mysql3.NewDeleteComorbidityRepository(db, kitlogger)
	deleteComorbidityService := service3.NewDeleteComorbidityService(deleteComorbidityRepo, kitlogger)
	deleteComorbidityEndpoint := handler3.MakeDeleteComorbidityEndpoint(deleteComorbidityService)
	deleteComorbidityEndpoint = handler3.DeleteComorbidityTransportMiddleware(kitlogger)(deleteComorbidityEndpoint)
	deleteComorbidityHandler := handler3.NewDeleteComorbidityHandler(config.GetString("paths.deleteComorbidity"), deleteComorbidityEndpoint)
	//////////////////////DELETE CONMORBILITY////////////////////////////////////////////////

	//////////////////////DELETE SYMPTOM////////////////////////////////////////////////
	deleteSymptomRepo := mysql5.NewDeleteSymptomRepository(db, kitlogger)
	deleteSymptomService := service5.NewDeleteSymptomService(deleteSymptomRepo, kitlogger)
	deleteSymptomEndpoint := handler5.MakeDeleteSymptomEndpoint(deleteSymptomService)
	deleteSymptomEndpoint = handler5.DeleteSymptomTransportMiddleware(kitlogger)(deleteSymptomEndpoint)
	deleteSymptomHandler := handler5.NewDeleteSymptomHandler(config.GetString("paths.deleteSymptom"), deleteSymptomEndpoint)
	//////////////////////DELETE SYMPTOM////////////////////////////////////////////////

	mux.Handle(config.GetString("paths.createComorbidity"), createComorbidityHandler)
	mux.Handle(config.GetString("paths.createSymptom"), createSymptomHandler)
	mux.Handle(config.GetString("paths.createInfoPatient"), createInfoPatientHandler)

	mux.Handle(config.GetString("paths.updateComorbidity"), updateComorbidityHandler)
	mux.Handle(config.GetString("paths.updateSymptom"), updateSymptomHandler)

	mux.Handle(config.GetString("paths.deleteComorbidity"), deleteComorbidityHandler)
	mux.Handle(config.GetString("paths.deleteSymptom"), deleteSymptomHandler)
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
