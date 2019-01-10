package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetDistricts controller to get all province entities without paging
var GetDistricts = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetAllDistricts())
}

// GetDistrictByProvinceID controller to get district by province id
var GetDistrictByProvinceID = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceID, err := strconv.Atoi(vars["provinceId"])
	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch district by province id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetDistrictByProvinceID(provinceID))
}
