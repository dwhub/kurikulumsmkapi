package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggProvinceResp HTTP status code and province model in message
// swagger:response provinceResp
type SwaggProvinceResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Province model
		Message []models.Province `json:"message"`
	}
}
