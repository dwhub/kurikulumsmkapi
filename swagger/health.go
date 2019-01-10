package swagger

import (
	"github.com/dwhub/kurikulumsmkapi/models"
)

// SwaggHealthResp HTTP status code and system model in message
// swagger:response healthResp
type SwaggHealthResp struct {
	// in:body
	Body struct {
		// HTTP status code
		Code int `json:"code"`
		// System model
		Message models.System `json:"message"`
	}
}
