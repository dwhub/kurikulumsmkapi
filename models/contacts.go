package models

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// Contact model and json tag
type Contact struct {
	ID         int    `json:"id"`
	DistrictID int    `json:"district_id"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

// ContactPaging contact model with paging
type ContactPaging struct {
	Contacts []Contact `json:"contacts"`
	Paging   Paging    `json:"paging"`
}

var contactHeaderQuery = "SELECT a.id_contact, a.id_kabupaten, b.kabupaten, a.nama_contact, a.no_telp, a.alamat "

var contactBaseQuery = contactHeaderQuery + `FROM tbl_contact a
						LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten `

var contactWithProvinceQuery = contactHeaderQuery + `FROM tbl_contact a
						LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten 
						LEFT JOIN tbl_provinsi c on b.id_provinsi = c.id_provinsi `

var contactPagingQuery = `LIMIT ? OFFSET ?`

var contactGetTotalRowsQuery = `SELECT COUNT(id_contact) FROM tbl_contact `

var contactWithProvinceGetTotalQuery = `SELECT COUNT(id_contact)
						FROM tbl_contact a
						LEFT JOIN tbl_kabupaten b on a.id_kabupaten = b.id_kabupaten 
						LEFT JOIN tbl_provinsi c on b.id_provinsi = c.id_provinsi `

// GetContacts fetch contacts with paging
func GetContacts(page int, pageSize int, provinceID int, districtID int) map[string]interface{} {
	var (
		contact       Contact
		contactPaging ContactPaging
	)

	var row *sql.Row

	if districtID > 0 {
		row = db.QueryRow(contactGetTotalRowsQuery+"where id_kabupaten = ?", districtID)
	} else {
		if provinceID > 0 {
			row = db.QueryRow(contactWithProvinceGetTotalQuery+"where c.id_provinsi = ?", provinceID)
		} else {
			row = db.QueryRow(contactGetTotalRowsQuery)
		}
	}

	var totalRow int

	err := row.Scan(&totalRow)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch contact status")

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
		rows, err = db.Query(contactBaseQuery+"where a.id_kabupaten = ? "+contactPagingQuery, districtID, pageSize, offset)
	} else {
		if provinceID > 0 {
			rows, err = db.Query(contactWithProvinceQuery+"where c.id_provinsi = ? "+contactPagingQuery, provinceID, pageSize, offset)
		} else {
			rows, err = db.Query(contactBaseQuery+contactPagingQuery, pageSize, offset)
		}
	}

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch contact status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&contact.ID, &contact.DistrictID, &contact.District, &contact.Name, &contact.Phone, &contact.Address)
		contactPaging.Contacts = append(contactPaging.Contacts, contact)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	contactPaging.Paging.Page = page
	contactPaging.Paging.Size = pageSize
	contactPaging.Paging.Total = totalPage

	return u.Message(http.StatusOK, contactPaging)
}

// GetAllContacts fetch all contacts without paging
func GetAllContacts() map[string]interface{} {
	var (
		contact  Contact
		contacts []Contact
	)

	rows, err := db.Query(contactBaseQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch contact status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&contact.ID, &contact.DistrictID, &contact.District, &contact.Name, &contact.Phone, &contact.Address)
		contacts = append(contacts, contact)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, contacts)
}
