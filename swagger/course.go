package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggCourseDurationResp HTTP status code and course duration model in message
// swagger:response courseDurationResp
type SwaggCourseDurationResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Province model
		Message []models.District `json:"message"`
	}
}
