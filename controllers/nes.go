package controllers

import (
	"net/http"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetNationalEducationStandards controller to get all national education standard entities without paging
var GetNationalEducationStandards = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetNationalEducationStandards())
}
