package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dwhub/kurikulumsmkapi/app"
	"github.com/dwhub/kurikulumsmkapi/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("Initializing Logging")
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Info("Initializing Environment")
	e := godotenv.Load()
	if e != nil {
		log.WithFields(log.Fields{
			"error": e,
		}).Error("Failed to initialize environment")
	}
}

func main() {

	log.Info("Kurikulum SMK Service is starting!")

	router := app.GetRouter()

	sw := http.FileServer(http.Dir("./explorer/"))
	router.PathPrefix("/explorer/").Handler(http.StripPrefix("/explorer/", sw))

	fs := http.FileServer(http.Dir("./files/"))
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("db_user"), os.Getenv("db_pass"), os.Getenv("db_host"), os.Getenv("db_name"))
	log.Info(connectionString)
	models.InitDB(connectionString)

	log.WithFields(log.Fields{
		"system": "startup",
	}).Info("Service is run at port " + port)

	err := http.ListenAndServe(":"+port, handlers.RecoveryHandler()(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to run service")
	}
}
