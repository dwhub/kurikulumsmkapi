package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// Province model and json tag
type Province struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

var provinceBaseQuery = `SELECT id_provinsi, provinsi, urutan_provinsi 
						FROM tbl_provinsi`

// GetAllProvinces fetch all provinces without paging
func GetAllProvinces() map[string]interface{} {
	var (
		province  Province
		provinces []Province
	)

	provinces = []Province{}

	rows, err := db.Query(provinceBaseQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch province status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&province.ID, &province.Name, &province.Order)
		provinces = append(provinces, province)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, provinces)
}
