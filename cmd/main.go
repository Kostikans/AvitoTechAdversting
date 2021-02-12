//  Golang service API for Avito
//
//  Swagger spec.
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//	- application/json
//  swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	apiMiddleware "github.com/Kostikans/AvitoTechadvertising/internal/app/middleware"

	"github.com/go-playground/validator/v10"

	advertisingDelivery "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/delivery/http"
	advertisingRepository "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/repository"
	advertisingUsecase "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/usecase"

	"github.com/Kostikans/AvitoTechadvertising/configs"
	"github.com/Kostikans/AvitoTechadvertising/internal/package/logger"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger.yaml"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

	return router
}

func InitDB() *sqlx.DB {
	var connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configs.BdConfig.User,
		configs.BdConfig.Password,
		configs.BdConfig.DBName)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {

	err := godotenv.Load("vars.env")
	if err != nil {
		log.Fatal(err)
	}
	configs.Init()
	db := InitDB()

	err = configs.ExportConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := logger.NewLogger(os.Stdout)

	r := NewRouter()
	r.Use(apiMiddleware.NewPanicMiddleware())
	r.Use(apiMiddleware.LoggerMiddleware(log))
	advRepo := advertisingRepository.NewAdvertisingRepository(db)
	validation := validator.New()
	advUsecase := advertisingUsecase.NewAdvertisingUsecase(advRepo, validation)
	advertisingDelivery.NewAdvertisingDelivery(r, advUsecase, log)

	if err := http.ListenAndServe(viper.GetString(configs.ConfigFields.AvitoServicePort), r); err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", viper.GetString(configs.ConfigFields.AvitoServicePort))
}
