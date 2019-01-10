package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggDistrictsResp HTTP status code and districts model in message
// swagger:response districtsResp
type SwaggDistrictsResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Province model
		Message []models.District `json:"message"`
	}
}
