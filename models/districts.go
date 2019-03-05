package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// District model and json tag
type District struct {
	ID         int    `json:"id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Name       string `json:"name"`
}

var districtBaseQuery = `SELECT a.id_kabupaten, a.id_provinsi, b.provinsi, a.kabupaten 
						FROM tbl_kabupaten a
						LEFT JOIN tbl_provinsi b on a.id_provinsi = b.id_provinsi `

// GetAllDistricts fetch all districts without paging
func GetAllDistricts() map[string]interface{} {
	var (
		district  District
		districts []District
	)

	districts = []District{}

	rows, err := db.Query(districtBaseQuery + " ORDER BY a.kabupaten ")

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch district status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&district.ID, &district.ProvinceID, &district.Province, &district.Name)
		districts = append(districts, district)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, districts)
}

// GetDistrictByProvinceID fetch district by ID
func GetDistrictByProvinceID(provinceID int) map[string]interface{} {
	var (
		district  District
		districts []District
	)

	districts = []District{}

	rows, err := db.Query(districtBaseQuery+"where a.id_provinsi = ?"+" ORDER BY a.kabupaten ", provinceID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch district status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&district.ID, &district.ProvinceID, &district.Province, &district.Name)
		districts = append(districts, district)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, districts)
}
