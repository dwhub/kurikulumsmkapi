package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// CourseAllocation model and json tag
type CourseAllocation struct {
	ID             int    `json:"id"`
	Order          int    `json:"order"`
	Group          int    `json:"group"`
	CompetencyID   int    `json:"competency_id"`
	Name           string `json:"name"`
	TimeAllocation int    `json:"time_allocation"`
}

var courseAllocationBaseQuery = `SELECT DISTINCT a.id_mapel, a.urutan_mapel, a.id_grup, b.id_kompetensi, a.nama_mapel, b.alokasi_waktu
									FROM tbl_mapel_alokasi b
									LEFT JOIN tbl_mapel a on b.id_mapel = a.id_mapel `

// GetCourseAllocations fetch all course allocation based on group without paging
func GetCourseAllocations(competencyID int, groupID int) map[string]interface{} {
	var (
		courseAllocation  CourseAllocation
		courseAllocations []CourseAllocation
	)

	courseAllocations = []CourseAllocation{}

	rows, err := db.Query(courseAllocationBaseQuery+" WHERE a.id_grup = ? AND b.id_kompetensi = ? ", groupID, competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch course duration status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&courseAllocation.ID, &courseAllocation.Order, &courseAllocation.Group, &courseAllocation.CompetencyID,
			&courseAllocation.Name, &courseAllocation.TimeAllocation)
		courseAllocations = append(courseAllocations, courseAllocation)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, courseAllocations)
}
