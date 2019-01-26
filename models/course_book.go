package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// CourseBook model and json tag
type CourseBook struct {
	ID           int    `json:"id"`
	CourseID     int    `json:"course_id"`
	Name         string `json:"name"`
	CompetencyID int    `json:"competency_id"`
	X            string `json:"x"`
	XI           string `json:"xi"`
	XII          string `json:"xii"`
	XIII         string `json:"xiii"`
	StudentBook  int    `json:"student_book"`
	TeacherBook  int    `json:"teacher_book"`
	Total        int    `json:"total"`
}

// Course model and json tag
type Course struct {
	ID      int    `json:"id"`
	Order   int    `json:"order"`
	GroupID int    `json:"group_id"`
	Name    string `json:"name"`
}

var courseBaseQuery = `SELECT a.id_mapel, a.id_grup, a.urutan_mapel, a.nama_mapel
						FROM tbl_mapel a `

var courseBookBaseQuery = `SELECT a.id_buku, a.id_mapel, b.nama_mapel, a.id_kompetensi, a.siswa_10 as X, a.siswa_11 as XI, a.siswa_12 as XII, a.siswa_13 as XIII,
								a.buku_siswa, a.buku_guru, a.buku_siswa + a.buku_guru as total
								FROM tbl_buku a
								LEFT JOIN tbl_mapel b ON b.id_mapel = a.id_mapel `

// GetCourses fetch all course based on competency without paging
func GetCourses(competencyID int) map[string]interface{} {
	var (
		course  Course
		courses []Course
	)

	courses = []Course{}

	rows, err := db.Query(courseBaseQuery+" WHERE a.id_kompetensi IN (0, ?) AND parent_id_mapel = 0 ", competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch courses status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&course.ID, &course.GroupID, &course.Order, &course.Name)
		courses = append(courses, course)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, courses)
}

// GetCourseBooks fetch all course book based on group and competency without paging
func GetCourseBooks(competencyID int, groupID int) map[string]interface{} {
	var (
		courseBook  CourseBook
		courseBooks []CourseBook
	)

	courseBooks = []CourseBook{}

	rows, err := db.Query(courseBookBaseQuery+" WHERE b.id_grup = ? AND a.id_kompetensi = ? ", groupID, competencyID)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch course book status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&courseBook.ID, &courseBook.CourseID, &courseBook.Name, &courseBook.CompetencyID,
			&courseBook.X, &courseBook.XI, &courseBook.XII, &courseBook.XIII, &courseBook.StudentBook,
			&courseBook.TeacherBook, &courseBook.Total)
		courseBooks = append(courseBooks, courseBook)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, courseBooks)
}
