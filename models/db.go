package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // used for mysql access
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

// Paging fields
type Paging struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

// InitDB initialize db access
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Failed to validate database connection string")
	}

	PingDB()
}

// PingDB check db connection
func PingDB() bool {
	var result bool

	err := db.Ping()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Failed to verify connection to database")
	}

	if err == nil {
		result = true
	}

	return result
}
