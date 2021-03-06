package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggGetSchoolsParam page param
// swagger:parameters listSchools
type SwaggGetSchoolsParam struct {
	// schools page param
	//
	// items.items.items.pattern: ^\d+$
	// in: query
	Page string `json:"page"`
	// schools page size param
	//
	// items.items.items.pattern: ^\d+$
	// in: query
	PageSize string `json:"pageSize"`
	// schools district id param
	//
	// in: query
	DistrictID string `json:"districtID"`
	// schools province id param
	//
	// in: query
	ProvinceID string `json:"provinceID"`
	// schools competency id param
	//
	// in: query
	CompetencyID string `json:"competencyID"`
	// schools type param
	//
	// in: query
	SchoolType string `json:"schoolType"`
	// Sub District type param
	//
	// in: query
	SubDistrict string `json:"subDistrict"`
}

// SwaggSchoolResp HTTP status code and school model in message
// swagger:response schoolResp
type SwaggSchoolResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Contact model
		Message []models.School `json:"message"`
	}
}

// SwaggSchoolPagingResp HTTP status code and school paging model in message
// swagger:response schoolPagingResp
type SwaggSchoolPagingResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// School model
		Message models.SchoolPaging `json:"message"`
	}
}
