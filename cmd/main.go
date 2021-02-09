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
	logOutput, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = configs.ExportConfig()
	if err != nil {
		log.Fatal(err)
	}
	defer logOutput.Close()
	log := logger.NewLogger(logOutput)
	if err != nil {
		log.Error(err)
	}

	r := NewRouter()

	err = http.ListenAndServe(viper.GetString(configs.ConfigFields.AvitoServicePort), r)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", viper.GetString(configs.ConfigFields.AvitoServicePort))
}
