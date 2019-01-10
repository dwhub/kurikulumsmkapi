package controllers

import (
	"net/http"

	"github.com/dwhub/kurikulumsmkapi/models"
	u "github.com/dwhub/kurikulumsmkapi/utils"
)

// CheckHealth controller to get the status of the server
var CheckHealth = func(w http.ResponseWriter, r *http.Request) {
	sys := models.System{}

	u.Respond(w, sys.CheckStatus())
}
