package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggNESResp HTTP status code and national education standard model in message
// swagger:response nesResp
type SwaggNESResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Expertise fields model
		Message []models.NationalEducationStandard `json:"message"`
	}
}
