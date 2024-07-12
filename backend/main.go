package main

import (
	"backend/db"
	"backend/router"
	"log"
	"net/http"
	"os"

	_ "backend/docs"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	var databaseURI string

	pflag.StringVar(&databaseURI, "databaseURI", "", "Database URI")
	pflag.Parse()

	viper.BindPFlag("database_uri", pflag.Lookup("databaseURI"))
	viper.AutomaticEnv()

	db.InitCluster(databaseURI)
	defer db.CloseCluster()

	r := router.Router()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
