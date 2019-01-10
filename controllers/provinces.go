package controllers

import (
	"net/http"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// GetAllProvinces controller to get all province entities without paging
var GetAllProvinces = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, models.GetAllProvinces())
}
