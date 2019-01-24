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

// SwaggCourseKIKDResp HTTP status code and course KI and KD model in message
// swagger:response courseKIKDResp
type SwaggCourseKIKDResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Course KI and KD model
		Message []models.CourseKIKD `json:"message"`
	}
}

// SwaggKIKDDetailResp HTTP status code and KI and KD detail model in message
// swagger:response KIKDDetailResp
type SwaggKIKDDetailResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// KI and KD detail model
		Message []models.KI `json:"message"`
	}
}
