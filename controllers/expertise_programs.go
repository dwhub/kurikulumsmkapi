package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetExpertisePrograms controller to get all expertise program entities without paging
var GetExpertisePrograms = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetExpertisePrograms())
}

// GetExpertiseProgramsByFieldID controller to get district by field id
var GetExpertiseProgramsByFieldID = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fieldID, err := strconv.Atoi(vars["fieldId"])
	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch expertise program by field id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetExpertiseProgramsByFieldID(fieldID))
}
