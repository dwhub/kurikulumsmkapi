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
		// Course Duration model
		Message []models.CourseDuration `json:"message"`
	}
}

// SwaggCourseAllocationResp HTTP status code and course allocation model in message
// swagger:response courseAllocationResp
type SwaggCourseAllocationResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Course ALlocation model
		Message []models.CourseAllocation `json:"message"`
	}
}
