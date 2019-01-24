package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// CourseKIKD model and json tag
type CourseKIKD struct {
	ID           int    `json:"id"`
	Order        int    `json:"order"`
	Group        int    `json:"group"`
	CompetencyID int    `json:"competency_id"`
	Name         string `json:"name"`
	ParentID     int    `json:"parent_id"`
	HaveChildren bool   `json:"have_children"`
}

// KI model and json tag
type KI struct {
	ID      int    `json:"id"`
	Order   string `json:"order"`
	Name    string `json:"name"`
	Details []KD   `json:"kds"`
}

// KD model and json tag
type KD struct {
	ID    int    `json:"id"`
	Order string `json:"order"`
	Name  string `json:"name"`
}

var getCourseKIKDBaseQuery = `SELECT a.id_mapel, a.urutan_mapel, a.parent_id_mapel, a.id_grup, a.id_kompetensi, 
									a.nama_mapel, a.has_children
									FROM tbl_mapel a `

var getCourseKIBaseQuery = `SELECT id_ki, uraian_ki, urutan_ki
							FROM smk.tbl_ki `

var getCourseKDBaseQuery = `SELECT id_kd, uraian_kd, urutan_kd
							FROM smk.tbl_kd `

// GetCourseKIKD fetch all course KI and KD based on competency and group without paging
func GetCourseKIKD(competencyID int, groupID int) map[string]interface{} {
	var (
		courseKIKD  CourseKIKD
		courseKIKDs []CourseKIKD
	)

	var finalKIKDs = []CourseKIKD{}

	rows, err := db.Query(getCourseKIKDBaseQuery+" WHERE a.id_grup = ? AND a.id_kompetensi = ? ", groupID, competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch course KI and KD status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&courseKIKD.ID, &courseKIKD.Order, &courseKIKD.ParentID, &courseKIKD.Group, &courseKIKD.CompetencyID,
			&courseKIKD.Name, &courseKIKD.HaveChildren)
		courseKIKDs = append(courseKIKDs, courseKIKD)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	for _, item := range courseKIKDs {

		if item.ParentID == 0 {
			finalKIKDs = append(finalKIKDs, item)
		}

		if item.ParentID == 0 && item.HaveChildren {
			for _, child := range courseKIKDs {
				if child.ParentID == item.ID {
					child.Name = "- " + child.Name
					finalKIKDs = append(finalKIKDs, child)
				}
			}
		}
	}

	return u.Message(http.StatusOK, finalKIKDs)
}

// GetKIKDDetail fetch all KI and KD detail based on competency and courseid without paging
func GetKIKDDetail(courseID int, competencyID int) map[string]interface{} {
	var (
		ki  KI
		kis []KI
	)

	var fixKis = []KI{}

	rows, err := db.Query(getCourseKIBaseQuery+" WHERE id_mapel = ? and id_kompetensi = ? ", courseID, competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch KI detail status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&ki.ID, &ki.Name, &ki.Order)
		kis = append(kis, ki)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	for _, item := range kis {
		var kd KD

		details, err := db.Query(getCourseKDBaseQuery+" WHERE id_ki = ? ", item.ID)

		if err != nil {
			log.WithFields(log.Fields{
				"status": "Failed",
				"error":  err,
			}).Info("Fetch KD detail status")

			return u.Message(http.StatusInternalServerError, "")
		}

		for details.Next() {
			err = details.Scan(&kd.ID, &kd.Name, &kd.Order)
			kd.Order = item.Order + "." + kd.Order
			item.Details = append(item.Details, kd)
			if err != nil {
				fmt.Print(err.Error())
			}
		}

		fixKis = append(fixKis, item)
	}

	return u.Message(http.StatusOK, fixKis)
}
