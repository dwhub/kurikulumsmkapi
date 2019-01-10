package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggExpertiseCompetenciesResp HTTP status code and expertise competencies model in message
// swagger:response expertiseCompetenciesResp
type SwaggExpertiseCompetenciesResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Expertise fields model
		Message []models.ExpertiseCompetency `json:"message"`
	}
}
