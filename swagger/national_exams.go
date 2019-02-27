package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggNationalExamsResp HTTP status code and national exam model in message
// swagger:response nationalExamsResp
type SwaggNationalExamsResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Expertise fields model
		Message []models.NationalExamStructure `json:"message"`
	}
}
