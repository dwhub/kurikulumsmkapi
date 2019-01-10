package controllers

import (
	"net/http"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetExpertiseFields controller to get all expertise field entities without paging
var GetExpertiseFields = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetExpertiseFields())
}

// GetCurriculumStructures controller to get expertise structure entities without paging
var GetCurriculumStructures = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetCurriculumStructures())
}
