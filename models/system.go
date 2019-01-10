package models

import (
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// System model and json tag
type System struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

// CheckStatus do some system and db check
func (system System) CheckStatus() map[string]interface{} {
	pingDB := PingDB()
	var httpStatus int

	if !pingDB {
		httpStatus = http.StatusInternalServerError
		system.Status = "DB Error"
		system.Database = "Database connection failed, please check log for more detailed information"
	} else {
		httpStatus = http.StatusOK
		system.Status = "OK"
		system.Database = "Database is up and running"
	}

	log.WithFields(log.Fields{
		"status": system.Status,
	}).Info("Server Status ")

	resp := u.Message(httpStatus, "success")
	resp["message"] = system

	return resp
}
