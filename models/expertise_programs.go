package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// ExpertiseProgram model and json tag
type ExpertiseProgram struct {
	ID               int    `json:"id"`
	ExpertiseFieldID string `json:"expertis_field_id"`
	ExpertiseField   string `json:"expertise_field"`
	Name             string `json:"name"`
	Order            int    `json:"order"`
}

var expertiseProgramBaseQuery = `SELECT a.id_program, a.id_bidang, b.bidang_keahlian, a.program_keahlian, a.urutan_program 
						FROM tbl_program_keahlian a
						LEFT JOIN tbl_bidang_keahlian b on a.id_bidang = b.id_bidang `

// GetExpertisePrograms fetch all expertise program without paging
func GetExpertisePrograms() map[string]interface{} {
	var (
		expertiseProgram  ExpertiseProgram
		expertisePrograms []ExpertiseProgram
	)

	expertisePrograms = []ExpertiseProgram{}

	rows, err := db.Query(expertiseProgramBaseQuery + " ORDER BY a.urutan_program")

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch expertise program status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&expertiseProgram.ID, &expertiseProgram.ExpertiseFieldID,
			&expertiseProgram.ExpertiseField, &expertiseProgram.Name, &expertiseProgram.Order)
		expertisePrograms = append(expertisePrograms, expertiseProgram)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, expertisePrograms)
}

// GetExpertiseProgramsByFieldID fetch Expertise program by field ID
func GetExpertiseProgramsByFieldID(expertiseFieldID int) map[string]interface{} {
	var (
		expertiseProgram  ExpertiseProgram
		expertisePrograms []ExpertiseProgram
	)

	expertisePrograms = []ExpertiseProgram{}

	rows, err := db.Query(expertiseProgramBaseQuery+"WHERE a.id_bidang = ? ORDER BY a.urutan_program", expertiseFieldID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch expertise program status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&expertiseProgram.ID, &expertiseProgram.ExpertiseFieldID,
			&expertiseProgram.ExpertiseField, &expertiseProgram.Name, &expertiseProgram.Order)
		expertisePrograms = append(expertisePrograms, expertiseProgram)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, expertisePrograms)
}
