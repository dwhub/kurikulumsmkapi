package controllers

import (
	"net/http"
	"strconv"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// GetSchools controller to get the schools entities
var GetSchools = func(w http.ResponseWriter, r *http.Request) {
	var pageSize int
	var page int
	var districtID int
	var provinceID int
	var competencyID int
	var schoolType int
	var err error

	pageSizeParam := r.URL.Query().Get("pageSize")
	pageParam := r.URL.Query().Get("page")
	districtIDParam := r.URL.Query().Get("districtID")
	provinceIDParam := r.URL.Query().Get("provinceID")
	competencyIDParam := r.URL.Query().Get("competencyID")
	schoolTypeParam := r.URL.Query().Get("schoolType")

	pageSize, err = strconv.Atoi(pageSizeParam)
	page, err = strconv.Atoi(pageParam)
	if len(districtIDParam) > 0 {
		districtID, err = strconv.Atoi(districtIDParam)
	}
	if len(provinceIDParam) > 0 {
		provinceID, err = strconv.Atoi(provinceIDParam)
	}
	if len(competencyIDParam) > 0 {
		competencyID, err = strconv.Atoi(competencyIDParam)
	}
	if len(schoolTypeParam) > 0 {
		schoolType, err = strconv.Atoi(schoolTypeParam)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch school status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetSchools(page, pageSize, districtID, provinceID, competencyID, schoolType))
}

// GetAllSchools controller to get all schools entities without paging
var GetAllSchools = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetAllSchools())
}
