package controllers

import (
	"net/http"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetNationalExams controller to get all national exam entities without paging
var GetNationalExams = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetNationalExams())
}
