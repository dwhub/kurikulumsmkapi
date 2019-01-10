package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggExpertiseProgramsResp HTTP status code and expertise program model in message
// swagger:response expertiseProgramsResp
type SwaggExpertiseProgramsResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Expertise Program model
		Message []models.ExpertiseProgram `json:"message"`
	}
}
