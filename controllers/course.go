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

// GetCourseAllocations controller to get course allocation by competency and group id
var GetCourseAllocations = func(w http.ResponseWriter, r *http.Request) {
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

	u.Respond(w, models.GetCourseAllocations(competencyID, groupID))
}

// GetCourseKIKDs controller to get course KI and KD by competency and group id
var GetCourseKIKDs = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competencyID, err := strconv.Atoi(vars["competencyId"])
	groupID, err := strconv.Atoi(vars["groupId"])

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch course KI and KD by competency and group id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	if groupID == 1 {
		competencyID = 0
	} else if groupID == 2 {
		competencyID = 0
	}

	u.Respond(w, models.GetCourseKIKD(competencyID, groupID))
}

// GetKIKDDetails controller to get course KI and KD by competency and course id
var GetKIKDDetails = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competencyID, err := strconv.Atoi(vars["competencyId"])
	courseID, err := strconv.Atoi(vars["courseId"])

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch KI and KD detail by competency and course id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetKIKDDetail(courseID, competencyID))
}

// GetCourseBooks controller to get course books by competency and group id
var GetCourseBooks = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competencyID, err := strconv.Atoi(vars["competencyId"])
	groupID, err := strconv.Atoi(vars["groupId"])

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch course books by competency and group id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetCourseBooks(competencyID, groupID))
}

// GetCourses controller to get courses by competency
var GetCourses = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	competencyID, err := strconv.Atoi(vars["competencyId"])

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Bad Request",
			"error":  err,
		}).Info("Fetch course by competency id status")

		resp := u.Message(http.StatusBadRequest, "")
		resp["message"] = "Request param is not valid"

		u.Respond(w, resp)
		return
	}

	u.Respond(w, models.GetCourses(competencyID))
}
