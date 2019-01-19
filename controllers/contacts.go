package controllers

import (
	"net/http"
	"strconv"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// GetContacts controller to get the contact entities
var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	var pageSize int
	var page int
	var districtID int
	var provinceID int
	var err error

	pageSizeParam := r.URL.Query().Get("pageSize")
	pageParam := r.URL.Query().Get("page")
	districtIDParam := r.URL.Query().Get("districtID")
	provinceIDParam := r.URL.Query().Get("provinceID")

	pageSize, err = strconv.Atoi(pageSizeParam)
	page, err = strconv.Atoi(pageParam)
	if len(districtIDParam) > 0 {
		districtID, err = strconv.Atoi(districtIDParam)
	}
	if len(provinceIDParam) > 0 {
		provinceID, err = strconv.Atoi(provinceIDParam)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch contact status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetContacts(page, pageSize, provinceID, districtID))
}

// GetAllContacts controller to get all contact entities without paging
var GetAllContacts = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetAllContacts())
}
