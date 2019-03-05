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
	Province    string `json:"province"`
}

// SchoolPaging school model with paging
type SchoolPaging struct {
	Schools []School `json:"schools"`
	Paging  Paging   `json:"paging"`
}

var schoolBaseQuery = `SELECT a.id_sekolah, a.id_kabupaten, b.kabupaten, a.npsn, a.nama_sekolah, a.status,
						a.alamat_sekolah, a.kec_sekolah, a.no_telp_sekolah, a.no_fax_sekolah, c.provinsi
						FROM tbl_sekolah a
						LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten
						LEFT JOIN tbl_provinsi c on b.id_provinsi = c.id_provinsi `

var schoolFilterQuery = `SELECT DISTINCT a.id_sekolah, a.id_kabupaten, b.kabupaten, a.npsn, a.nama_sekolah, a.status,
							a.alamat_sekolah, a.kec_sekolah, a.no_telp_sekolah, a.no_fax_sekolah, c.provinsi
							FROM tbl_sekolah a
							LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten
							LEFT JOIN tbl_provinsi c on b.id_provinsi = c.id_provinsi
							LEFT JOIN tbl_kompetensi_sekolah d on a.id_sekolah = d.id_sekolah `

var schoolPagingQuery = ` LIMIT ? OFFSET ? `

var schoolGetTotalRowsQuery = `SELECT COUNT(id_sekolah) FROM tbl_sekolah `

var schoolFilterGetTotalRowsQuery = `SELECT COUNT(DISTINCT a.id_sekolah)
										FROM tbl_sekolah a 
										LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten
										LEFT JOIN tbl_provinsi c on b.id_provinsi = c.id_provinsi
										LEFT JOIN tbl_kompetensi_sekolah d on a.id_sekolah = d.id_sekolah `

// GetSchools fetch schools with paging
func GetSchools(page int, pageSize int, districtID int, provinceID int, competencyID int, schoolType int, subDistrict string) map[string]interface{} {
	var (
		school       School
		schoolPaging SchoolPaging
	)

	schoolPaging.Schools = []School{}

	var row *sql.Row

	var districtFilter = ""
	if districtID > 0 {
		districtFilter = " b.id_kabupaten = ? "
	}

	var provinceFilter = ""
	if provinceID > 0 {
		provinceFilter = " c.id_provinsi = ? "
	}

	var competencyFilter = ""
	if competencyID > 0 {
		competencyFilter = " d.id_kompetensi = ? "
	}

	var subDistrictFilter = ""
	if len(subDistrict) > 0 {
		subDistrictFilter = " a.kec_sekolah = ? "
	}

	var schoolTypeFilter = ""

	if schoolType == 1 {
		schoolTypeFilter = " AND a.status = 'N' "
	} else if schoolType == 2 {
		schoolTypeFilter = " AND a.status = 'S' "
	}

	if districtID == 0 && provinceID == 0 && competencyID == 0 && schoolType == 0 && len(subDistrict) == 0 {
		row = db.QueryRow(schoolFilterGetTotalRowsQuery)
	} else {
		if districtID > 0 && len(subDistrict) == 0 && competencyID == 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+districtFilter+schoolTypeFilter, districtID)
		} else if districtID == 0 && provinceID > 0 && len(subDistrict) == 0 && competencyID == 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+provinceFilter+schoolTypeFilter, provinceID)
		} else if districtID == 0 && provinceID == 0 && len(subDistrict) == 0 && competencyID > 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+competencyFilter+schoolTypeFilter, competencyID)
		} else if districtID == 0 && provinceID > 0 && len(subDistrict) == 0 && competencyID > 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+provinceFilter+" AND "+competencyFilter+schoolTypeFilter, provinceID, competencyID)
		} else if districtID > 0 && competencyID > 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+districtFilter+" AND "+competencyFilter+schoolTypeFilter, districtID, competencyID)
		} else if len(subDistrict) > 0 && competencyID == 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+subDistrictFilter+schoolTypeFilter, subDistrict)
		} else if len(subDistrict) > 0 && competencyID > 0 {
			row = db.QueryRow(schoolFilterGetTotalRowsQuery+"where "+subDistrictFilter+" AND "+competencyFilter+schoolTypeFilter, subDistrict, competencyID)
		}
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

	if districtID == 0 && provinceID == 0 && competencyID == 0 && schoolType == 0 && len(subDistrict) == 0 {
		rows, err = db.Query(schoolFilterQuery+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, pageSize, offset)
	} else {
		if districtID > 0 && len(subDistrict) == 0 && competencyID == 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+districtFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, districtID, pageSize, offset)
		} else if districtID == 0 && provinceID > 0 && len(subDistrict) == 0 && competencyID == 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+provinceFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, provinceID, pageSize, offset)
		} else if districtID == 0 && provinceID == 0 && len(subDistrict) == 0 && competencyID > 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+competencyFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, competencyID, pageSize, offset)
		} else if districtID == 0 && provinceID > 0 && len(subDistrict) == 0 && competencyID > 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+provinceFilter+" AND "+competencyFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, provinceID, competencyID, pageSize, offset)
		} else if districtID > 0 && competencyID > 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+districtFilter+" AND "+competencyFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, districtID, competencyID, pageSize, offset)
		} else if len(subDistrict) > 0 && competencyID == 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+subDistrictFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, subDistrict, pageSize, offset)
		} else if len(subDistrict) > 0 && competencyID > 0 {
			rows, err = db.Query(schoolFilterQuery+"where "+subDistrictFilter+" AND "+competencyFilter+schoolTypeFilter+" ORDER BY a.status, a.id_sekolah "+schoolPagingQuery, subDistrict, competencyID, pageSize, offset)
		}
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
			&school.Status, &school.Address, &school.SubDistrict, &school.Phone, &school.Fax, &school.Province)
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
			&school.Status, &school.Address, &school.SubDistrict, &school.Phone, &school.Fax, &school.Province)
		schools = append(schools, school)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, schools)
}
