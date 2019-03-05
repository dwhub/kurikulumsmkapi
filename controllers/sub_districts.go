package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetSubDistricts controller to get all sub district entities without paging
var GetSubDistricts = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetAllSubDistricts())
}

// GetSubDistrictByDistrictID controller to get sub district by district id
var GetSubDistrictByDistrictID = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	districtID, err := strconv.Atoi(vars["districtId"])
	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch sub district by district id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetSubDistrictByDistrictID(districtID))
}
