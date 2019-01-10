package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggExpertiseFieldsResp HTTP status code and expertise fields model in message
// swagger:response expertiseFieldResp
type SwaggExpertiseFieldsResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Expertise fields model
		Message []models.ExpertiseField `json:"message"`
	}
}

// SwaggCurriculumStructureResp HTTP status code and expertise fields model in message
// swagger:response curriculumStructureResp
type SwaggCurriculumStructureResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Curriculum structure fields model
		Message []models.ExpertiseStructure `json:"message"`
	}
}
