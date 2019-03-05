package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// SubDistrict model and json tag
type SubDistrict struct {
	DistrictID string `json:"district_id"`
	District   string `json:"district"`
	Name       string `json:"name"`
}

var subDistrictBaseQuery = `SELECT DISTINCT a.id_kabupaten, b.kabupaten, a.kec_sekolah
							FROM tbl_sekolah a
							LEFT JOIN tbl_kabupaten b ON a.id_kabupaten = b.id_kabupaten `

// GetAllSubDistricts fetch all districts without paging
func GetAllSubDistricts() map[string]interface{} {
	var (
		subDistrict  SubDistrict
		subDistricts []SubDistrict
	)

	subDistricts = []SubDistrict{}

	rows, err := db.Query(subDistrictBaseQuery + " ORDER BY id_kabupaten, kec_sekolah ")

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch sub district status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&subDistrict.DistrictID, &subDistrict.District, &subDistrict.Name)
		subDistricts = append(subDistricts, subDistrict)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, subDistricts)
}

// GetSubDistrictByDistrictID fetch district by ID
func GetSubDistrictByDistrictID(ditrictID int) map[string]interface{} {
	var (
		subDistrict  SubDistrict
		subDistricts []SubDistrict
	)

	subDistricts = []SubDistrict{}

	rows, err := db.Query(subDistrictBaseQuery+"WHERE a.id_kabupaten = ?"+" ORDER BY id_kabupaten, kec_sekolah ", ditrictID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch sub district status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&subDistrict.DistrictID, &subDistrict.District, &subDistrict.Name)
		subDistricts = append(subDistricts, subDistrict)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, subDistricts)
}
