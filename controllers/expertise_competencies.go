package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetExpertiseCompetencies controller to get all expertise competency entities without paging
var GetExpertiseCompetencies = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetExpertiseCompetencies())
}

// GetExpertiseCompetenciesByProgramID controller to get expertise competencies by program id
var GetExpertiseCompetenciesByProgramID = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	programID, err := strconv.Atoi(vars["programId"])
	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch expertise competency by program id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetExpertiseCompetenciesByProgramID(programID))
}
