package models

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// School model and json tag
type School struct {
	ID          int    `json:"id"`
	DistrictID  int    `json:"district_id"`
	District    string `json:"district"`
	NPSN        string `json:"npsn"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Address     string `json:"address"`
	SubDistrict string `json:"sub_district"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
}

// SchoolPaging school model with paging
type SchoolPaging struct {
	Schools []School `json:"schools"`
	Paging  Paging   `json:"paging"`
}

var schoolBaseQuery = `SELECT a.id_sekolah, a.id_kabupaten, b.kabupaten, a.npsn, a.nama_sekolah, a.status,
						a.alamat_sekolah, a.kec_sekolah, a.no_telp_sekolah, a.no_fax_sekolah
						FROM tbl_sekolah a
						LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten `

var schoolPagingQuery = `LIMIT ? OFFSET ?`

var schoolGetTotalRowsQuery = `SELECT COUNT(id_sekolah) FROM tbl_sekolah `

// GetSchools fetch schools with paging
func GetSchools(page int, pageSize int, districtID int) map[string]interface{} {
	var (
		school       School
		schoolPaging SchoolPaging
	)

	var row *sql.Row

	if districtID > 0 {
		row = db.QueryRow(schoolGetTotalRowsQuery+"where id_kabupaten = ?", districtID)
	} else {
		row = db.QueryRow(schoolGetTotalRowsQuery)
	}

	var totalRow int

	err := row.Scan(&totalRow)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch school status")

		return u.Message(http.StatusInternalServerError, "")
	}

	var offset int
	d := float64(totalRow) / float64(pageSize)
	totalPage := int(math.Ceil(d))

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}

	var rows *sql.Rows

	if districtID > 0 {
		rows, err = db.Query(schoolBaseQuery+"where a.id_kabupaten = ? "+schoolPagingQuery, districtID, pageSize, offset)
	} else {
		rows, err = db.Query(schoolBaseQuery+schoolPagingQuery, pageSize, offset)
	}

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch school status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&school.ID, &school.DistrictID, &school.District, &school.NPSN, &school.Name,
			&school.Status, &school.Address, &school.SubDistrict, &school.Phone, &school.Fax)
		schoolPaging.Schools = append(schoolPaging.Schools, school)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	schoolPaging.Paging.Page = page
	schoolPaging.Paging.Size = pageSize
	schoolPaging.Paging.Total = totalPage

	return u.Message(http.StatusOK, schoolPaging)
}

// GetAllSchools fetch all schools without paging
func GetAllSchools() map[string]interface{} {
	var (
		school  School
		schools []School
	)

	rows, err := db.Query(schoolBaseQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch school status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&school.ID, &school.DistrictID, &school.District, &school.NPSN, &school.Name,
			&school.Status, &school.Address, &school.SubDistrict, &school.Phone, &school.Fax)
		schools = append(schools, school)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, schools)
}
