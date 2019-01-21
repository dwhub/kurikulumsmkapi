package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// CourseDuration model and json tag
type CourseDuration struct {
	ID           int    `json:"id"`
	Order        int    `json:"order"`
	Group        int    `json:"group"`
	CompetencyID int    `json:"competency_id"`
	Name         string `json:"name"`
	X1           string `json:"x1"`
	X2           string `json:"x2"`
	XI1          string `json:"xi1"`
	XI2          string `json:"xi2"`
	XII1         string `json:"xii1"`
	XII2         string `json:"xii2"`
	XIII1        string `json:"xiii1"`
	XIII2        string `json:"xiii2"`
}

var courseDurationBaseQuery = `SELECT a.id_mapel, a.urutan_mapel, a.id_grup, a.id_kompetensi, a.nama_mapel, 
								b.x_1, b.x_2, b.xi_1, b.xi_2, b.xii_1, b.xii_2, b.xiii_1, b.xiii_2
								FROM tbl_mapel a
								LEFT JOIN tbl_mapel_matriks b on b.id_mapel = a.id_mapel AND b.id_kompetensi = a.id_kompetensi `

// GetCourseDurations fetch all course duration based on group without paging
func GetCourseDurations(competencyID int, groupID int) map[string]interface{} {
	var (
		courseDuration  CourseDuration
		courseDurations []CourseDuration
	)

	rows, err := db.Query(courseDurationBaseQuery+" WHERE a.id_grup = ? AND a.id_kompetensi = ? ", groupID, competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch course duration status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&courseDuration.ID, &courseDuration.Order, &courseDuration.Group, &courseDuration.CompetencyID,
			&courseDuration.Name, &courseDuration.X1, &courseDuration.X2, &courseDuration.XI1, &courseDuration.XI2,
			&courseDuration.XII1, &courseDuration.XII2, &courseDuration.XIII1, &courseDuration.XIII2)
		courseDurations = append(courseDurations, courseDuration)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, courseDurations)
}
