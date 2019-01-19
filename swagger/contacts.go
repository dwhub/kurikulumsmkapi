package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggGetContactParam page param
// swagger:parameters listContacts
type SwaggGetContactParam struct {
	// contacts page param
	//
	// items.items.items.pattern: ^\d+$
	// in: query
	Page string `json:"page"`
	// contacts page size param
	//
	// items.items.items.pattern: ^\d+$
	// in: query
	PageSize string `json:"pageSize"`
	// contacts province id param
	//
	// in: query
	ProvinceID string `json:"provinceID"`
	// contacts district id param
	//
	// in: query
	DistrictID string `json:"districtID"`
}

// SwaggContactResp HTTP status code and contact model in message
// swagger:response contactResp
type SwaggContactResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Contact model
		Message []models.Contact `json:"message"`
	}
}

// SwaggContactPagingResp HTTP status code and contact paging model in message
// swagger:response contactPagingResp
type SwaggContactPagingResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// Contact model
		Message models.ContactPaging `json:"message"`
	}
}
