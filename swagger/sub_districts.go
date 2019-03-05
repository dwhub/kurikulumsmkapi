package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggSubDistrictsResp HTTP status code and sub districts model in message
// swagger:response subDistrictsResp
type SwaggSubDistrictsResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Province model
		Message []models.SubDistrict `json:"message"`
	}
}
