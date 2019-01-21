package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetCourseDurations controller to get course duration by competency and group id
var GetCourseDurations = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competencyID, err := strconv.Atoi(vars["competencyId"])
	groupID, err := strconv.Atoi(vars["groupId"])

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch course duration by competency and group id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetCourseDurations(competencyID, groupID))
}
