package models

import (
	"fmt"
	"net/http"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// ExpertiseField model and json tag
type ExpertiseField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// ExpertiseStructure data as composite
type ExpertiseStructure struct {
	ID       int                  `json:"id"`
	ParentID int                  `json:"parent_id"`
	Name     string               `json:"name"`
	Children []ExpertiseStructure `json:"children"`
}

// FlatExpertiseStructure data as flat data
type FlatExpertiseStructure struct {
	FieldID         int    `json:"field_id"`
	ProgramID       int    `json:"program_id"`
	CompetencyID    int    `json:"competency_id"`
	FieldTitle      string `json:"name"`
	ProgramTitle    string `json:"program_title"`
	CompetencyTitle string `json:"competency_title"`
}

var efieldBaseQuery = `SELECT id_bidang, bidang_keahlian, urutan_bidang 
						FROM tbl_bidang_keahlian
						ORDER BY urutan_bidang`

var structCurriculumQuery = `SELECT a.id_bidang, b.id_program, c.id_kompetensi, 
							CONCAT(a.urutan_bidang, '. ', a.bidang_keahlian) as title_bidang, 
							CONCAT(a.urutan_bidang, '.', b.urutan_program, '. ', b.program_keahlian) as title_program, 
							CONCAT(a.urutan_bidang, '.', b.urutan_program, '.', c.urutan_kompetensi, '. ', c.kompetensi_keahlian) as title_kompetensi
							FROM tbl_bidang_keahlian a 
							LEFT JOIN tbl_program_keahlian b on a.id_bidang = b.id_bidang
							LEFT JOIN tbl_kompetensi_keahlian c on b.id_program = c.id_program`

// GetExpertiseFields fetch all provinces without paging
func GetExpertiseFields() map[string]interface{} {
	var (
		expertiseField  ExpertiseField
		expertiseFields []ExpertiseField
	)

	expertiseFields = []ExpertiseField{}

	rows, err := db.Query(efieldBaseQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch competency field status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&expertiseField.ID, &expertiseField.Name, &expertiseField.Order)
		expertiseFields = append(expertiseFields, expertiseField)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, expertiseFields)
}

// GetCurriculumStructures fetch curriculum structure
func GetCurriculumStructures() map[string]interface{} {
	var (
		flatExpertiseStructure  FlatExpertiseStructure
		flatExpertiseStructures []FlatExpertiseStructure
	)

	rows, err := db.Query(structCurriculumQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch structure curriculum status")

		return u.Message(http.StatusInternalServerError, "")
	}

	for rows.Next() {
		err = rows.Scan(&flatExpertiseStructure.FieldID,
			&flatExpertiseStructure.ProgramID,
			&flatExpertiseStructure.CompetencyID,
			&flatExpertiseStructure.FieldTitle,
			&flatExpertiseStructure.ProgramTitle,
			&flatExpertiseStructure.CompetencyTitle,
		)
		flatExpertiseStructures = append(flatExpertiseStructures, flatExpertiseStructure)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return u.Message(http.StatusOK, transformStructurCurriculum(flatExpertiseStructures))
}

// transformStructurCurriculum transform flat data to composite structure curriculum
func transformStructurCurriculum(flatExpertiseStructure []FlatExpertiseStructure) []ExpertiseStructure {
	var (
		field        ExpertiseStructure
		program      ExpertiseStructure
		competency   ExpertiseStructure
		fields       []ExpertiseStructure
		programs     []ExpertiseStructure
		competencies []ExpertiseStructure
		result       []ExpertiseStructure
	)

	if flatExpertiseStructure != nil {
		for _, item := range flatExpertiseStructure {
			if !itemAlreadyExist(fields, item.FieldID) {
				field.ID = item.FieldID
				//field.ChildID = item.ProgramID
				field.Name = item.FieldTitle

				fields = append(fields, field)
			}

			if !itemAlreadyExist(programs, item.ProgramID) {
				program.ID = item.ProgramID
				//program.ChildID = item.CompetencyID
				program.ParentID = item.FieldID
				program.Name = item.ProgramTitle

				programs = append(programs, program)
			}

			competency.ID = item.CompetencyID
			competency.ParentID = item.ProgramID
			competency.Name = item.CompetencyTitle

			competencies = append(competencies, competency)
		}

		result = mergeExpertiseStructure(fields, programs, competencies)
	}

	return result
}

func itemAlreadyExist(items []ExpertiseStructure, ID int) bool {
	for _, item := range items {
		if item.ID == ID {
			return true
		}
	}

	return false
}

func mergeExpertiseStructure(fields []ExpertiseStructure, programs []ExpertiseStructure, competencies []ExpertiseStructure) []ExpertiseStructure {
	var (
		field      ExpertiseStructure
		program    ExpertiseStructure
		competency ExpertiseStructure
		programRes []ExpertiseStructure
		result     []ExpertiseStructure
	)

	if programs != nil {
		for _, itemProg := range programs {
			program = ExpertiseStructure{}
			program.ID = itemProg.ID
			program.Name = itemProg.Name
			program.ParentID = itemProg.ParentID

			for _, itemComp := range competencies {
				if itemComp.ParentID == itemProg.ID {
					competency.ID = itemComp.ID
					competency.Name = itemComp.Name
					competency.ParentID = itemComp.ParentID

					program.Children = append(program.Children, itemComp)
				}
			}

			programRes = append(programRes, program)
		}
	}

	if fields != nil {
		for _, itemField := range fields {
			field = ExpertiseStructure{}
			field.ID = itemField.ID
			field.Name = itemField.Name

			for _, itemProg := range programRes {
				if itemProg.ParentID == itemField.ID {
					field.Children = append(field.Children, itemProg)
				}
			}

			result = append(result, field)
		}
	}

	return result
}
