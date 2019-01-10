package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// ExpertiseCompetency model and json tag
type ExpertiseCompetency struct {
	ID                 int    `json:"id"`
	ExpertiseProgramID string `json:"expertis_program_id"`
	ExpertiseProgram   string `json:"expertise_program"`
	Name               string `json:"name"`
	Duration           int    `json:"duration"`
	Order              int    `json:"order"`
}

var expertiseCompetencyBaseQuery = `SELECT a.id_kompetensi, a.id_program, b.program_keahlian, a.kompetensi_keahlian, a.waktu, a.urutan_kompetensi 
						FROM tbl_kompetensi_keahlian a
						LEFT JOIN tbl_program_keahlian b on a.id_program = b.id_program `

// GetExpertiseCompetencies fetch all expertise competency without paging
func GetExpertiseCompetencies() map[string]interface{} {
	var (
		expertiseCompetency   ExpertiseCompetency
		expertiseCompetencies []ExpertiseCompetency
	)

	rows, err := db.Query(expertiseCompetencyBaseQuery + " ORDER BY a.urutan_kompetensi")

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch expertise competency status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&expertiseCompetency.ID, &expertiseCompetency.ExpertiseProgramID,
			&expertiseCompetency.ExpertiseProgram, &expertiseCompetency.Name, &expertiseCompetency.Duration, &expertiseCompetency.Order)
		expertiseCompetencies = append(expertiseCompetencies, expertiseCompetency)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, expertiseCompetencies)
}

// GetExpertiseCompetenciesByProgramID fetch Expertise competency by program ID
func GetExpertiseCompetenciesByProgramID(expertiseProgramID int) map[string]interface{} {
	var (
		expertiseCompetency   ExpertiseCompetency
		expertiseCompetencies []ExpertiseCompetency
	)

	rows, err := db.Query(expertiseCompetencyBaseQuery+"WHERE a.id_program = ? ORDER BY a.urutan_kompetensi", expertiseProgramID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch expertise competency status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&expertiseCompetency.ID, &expertiseCompetency.ExpertiseProgramID,
			&expertiseCompetency.ExpertiseProgram, &expertiseCompetency.Name, &expertiseCompetency.Duration, &expertiseCompetency.Order)
		expertiseCompetencies = append(expertiseCompetencies, expertiseCompetency)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, expertiseCompetencies)
}
