package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// NationalExamStructure model
type NationalExamStructure struct {
	ID        int                     `json:"id"`
	ParentID  int                     `json:"parent_id"`
	Name      string                  `json:"name"`
	Spacing   int                     `json:"spacing"`
	PdfPath   string                  `json:"pdf_path"`
	CoverPath string                  `json:"cover_path"`
	Children  []NationalExamStructure `json:"children"`
}

// UNStandard model
type UNStandard struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Sub       string `json:"sub"`
	PdfPath   string `json:"pdf_path"`
	CoverPath string `json:"cover_path"`
}

// USBNStandard model
type USBNStandard struct {
	ParentID  int            `json:"parent_id"`
	Parent    string         `json:"parent"`
	ID        sql.NullInt64  `json:"id"`
	Name      sql.NullString `json:"name"`
	PdfPath   string         `json:"pdf_path"`
	CoverPath string         `json:"cover_path"`
}

// FlatNationalExamStruct data as flat data
type FlatNationalExamStruct struct {
	FieldID         int    `json:"field_id"`
	ProgramID       int    `json:"program_id"`
	CompetencyID    int    `json:"competency_id"`
	FieldTitle      string `json:"name"`
	ProgramTitle    string `json:"program_title"`
	CompetencyTitle string `json:"competency_title"`
	PdfPath         string `json:"pdf_path"`
	JpegCover       string `json:"jpeg_cover"`
}

var unStandardQuery = `SELECT a.id_mapel, b.nama_mapel, a.sub, a.pdf, a.cover
						FROM tbl_kisi_un_nasional a
						LEFT JOIN tbl_mapel b ON a.id_mapel = b.id_mapel`

var structCurriculumQueryWithPdf = `SELECT a.id_bidang, b.id_program, c.id_kompetensi, 
								CONCAT(a.urutan_bidang, '. ', a.bidang_keahlian) as title_bidang, 
								CONCAT(a.urutan_bidang, '.', b.urutan_program, '. ', b.program_keahlian) as title_program, 
								CONCAT(a.urutan_bidang, '.', b.urutan_program, '.', c.urutan_kompetensi, '. ', c.kompetensi_keahlian) as title_kompetensi,
								d.pdf, d.cover
								FROM tbl_bidang_keahlian a 
								LEFT JOIN tbl_program_keahlian b on a.id_bidang = b.id_bidang
								LEFT JOIN tbl_kompetensi_keahlian c on b.id_program = c.id_program
								LEFT JOIN tbl_kisi_un_kejuruan d on c.id_kompetensi = d.id_kompetensi`

var structCurriculumUSBNQueryWithPdf = `SELECT a.id_bidang, b.id_program, c.id_kompetensi, 
								CONCAT(a.urutan_bidang, '. ', a.bidang_keahlian) as title_bidang, 
								CONCAT(a.urutan_bidang, '.', b.urutan_program, '. ', b.program_keahlian) as title_program, 
								CONCAT(a.urutan_bidang, '.', b.urutan_program, '.', c.urutan_kompetensi, '. ', c.kompetensi_keahlian) as title_kompetensi,
								d.pdf, d.cover
								FROM tbl_bidang_keahlian a 
								LEFT JOIN tbl_program_keahlian b on a.id_bidang = b.id_bidang
								LEFT JOIN tbl_kompetensi_keahlian c on b.id_program = c.id_program
								LEFT JOIN tbl_kisi_usbn_kejuruan d on c.id_kompetensi = d.id_kompetensi`

var usbnStandardQuery = `SELECT a.id_mapel, a.nama_mapel as parent, b.id_mapel, b.nama_mapel, c.pdf, c.cover
							FROM tbl_mapel a
							LEFT JOIN tbl_mapel b ON a.id_mapel = b.parent_id_mapel
							RIGHT JOIN tbl_kisi_usbn_nasional c ON b.id_mapel = c.id_mapel OR a.id_mapel = c.id_mapel
							WHERE a.parent_id_mapel = 0`

// GetNationalExams fetch all national exam data without paging
func GetNationalExams() map[string]interface{} {
	var (
		schedule               NationalExamStructure
		nationalExamStructures []NationalExamStructure
	)

	nationalExamStructures = []NationalExamStructure{}

	// Add kisi kisi section
	kisiUNUSBN := LoadKisiUSBN()

	for _, item := range kisiUNUSBN {
		nationalExamStructures = append(nationalExamStructures, item)
	}
	// end kisi kisi section

	// Add jadwal item
	schedule.ID = 93
	schedule.Name = "Jadwal"

	nationalExamStructures = append(nationalExamStructures, schedule)
	// end jadwal item

	// Add tata tertib
	tataTertib := LoadTataTertib()

	for _, item := range tataTertib {
		nationalExamStructures = append(nationalExamStructures, item)
	}
	// end tata tertib

	// Add time allocation
	timeAllocations := LoadTimeAllocation()

	for _, item := range timeAllocations {
		nationalExamStructures = append(nationalExamStructures, item)
	}
	// end tata tertib

	return u.Message(http.StatusOK, nationalExamStructures)
}

// LoadKisiUSBN fetach kisi kisi UN USBN data
func LoadKisiUSBN() []NationalExamStructure {
	var (
		root         NationalExamStructure
		kisiUN       NationalExamStructure
		ujiKK        NationalExamStructure
		kisiUSBN     NationalExamStructure
		mn           NationalExamStructure
		usbnUjiKK    NationalExamStructure
		usbnMn       []NationalExamStructure
		usbnKeahlian []NationalExamStructure
		results      []NationalExamStructure
	)

	kisiUN.ID = 91
	kisiUN.Name = "Kisi-Kisi Ujian Nasional (UN) 2019"
	kisiUN.ParentID = 90
	kisiUN.Spacing = 1

	kisiUNStds := LoadKisiUNStd()

	for _, item := range kisiUNStds {
		kisiUN.Children = append(kisiUN.Children, item)
	}

	ujiKK.ID = 100
	ujiKK.Name = "Uji Kompetensi Keahlian"
	ujiKK.ParentID = 91
	ujiKK.Spacing = 2

	kisiUNKejuruan := loadKisiUNUSBNKejuruan(structCurriculumQueryWithPdf)

	for _, item := range kisiUNKejuruan {
		ujiKK.Children = append(ujiKK.Children, item)
	}

	kisiUN.Children = append(kisiUN.Children, ujiKK)

	kisiUSBN.ID = 92
	kisiUSBN.Name = "Kisi-Kisi Ujian Sekolah Berstandar Nasional (USBN) 2019"
	kisiUSBN.ParentID = 90
	kisiUSBN.Spacing = 1

	usbnMn = loadMuatanNasional()
	usbnKeahlian = loadKisiUNUSBNKejuruan(structCurriculumUSBNQueryWithPdf)

	mn.ID = 101
	mn.Name = "Muatan Nasional"
	mn.ParentID = 92
	mn.Spacing = 2

	for _, item := range usbnMn {
		mn.Children = append(mn.Children, item)
	}

	usbnUjiKK.ID = 102
	usbnUjiKK.Name = "Uji Kompetensi Keahlian"
	usbnUjiKK.ParentID = 92
	usbnUjiKK.Spacing = 2

	for _, item := range usbnKeahlian {
		usbnUjiKK.Children = append(usbnUjiKK.Children, item)
	}

	kisiUSBN.Children = append(kisiUSBN.Children, mn)
	kisiUSBN.Children = append(kisiUSBN.Children, usbnUjiKK)

	root.ID = 90
	root.Name = "Kisi-Kisi UN dan USBN 2019"
	root.Children = append(root.Children, kisiUN)
	root.Children = append(root.Children, kisiUSBN)

	results = append(results, root)

	return results
}

func loadMuatanNasional() []NationalExamStructure {
	var (
		usbnStandard USBNStandard
		queryResult  []USBNStandard
		parents      []NationalExamStructure
		childs       []NationalExamStructure
		results      []NationalExamStructure
	)

	rows, err := db.Query(usbnStandardQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch USBN standard status")

		return results
	}

	for rows.Next() {
		err = rows.Scan(&usbnStandard.ParentID, &usbnStandard.Parent, &usbnStandard.ID, &usbnStandard.Name, &usbnStandard.PdfPath, &usbnStandard.CoverPath)
		queryResult = append(queryResult, usbnStandard)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	for _, item := range queryResult {
		temp := NationalExamStructure{}
		parentExist := false

		for _, p := range parents {
			if item.ParentID == p.ID {
				temp = p
				parentExist = true
			}
		}

		if !item.ID.Valid {
			pdfPath := strings.Split(item.PdfPath, "/")
			coverPath := strings.Split(item.CoverPath, "/")

			temp.PdfPath = pdfPath[len(pdfPath)-1]
			temp.CoverPath = coverPath[len(coverPath)-1]
		}

		if item.ID.Valid {
			temp.ID = int(item.ID.Int64)
			temp.ParentID = item.ParentID
			temp.Name = item.Name.String
			temp.PdfPath = item.PdfPath
			temp.CoverPath = item.CoverPath

			childs = append(childs, temp)
		}

		if !parentExist {
			temp.ID = item.ParentID
			temp.Name = item.Parent

			parents = append(parents, temp)
		}
	}

	for _, parent := range parents {
		tmpParent := NationalExamStructure{}
		tmpParent.ID = parent.ID
		tmpParent.Name = parent.Name
		tmpParent.ParentID = parent.ParentID
		tmpParent.PdfPath = parent.PdfPath
		tmpParent.CoverPath = parent.CoverPath
		tmpParent.Spacing = 3

		for _, child := range childs {
			if child.ParentID == parent.ID {
				tmpChild := NationalExamStructure{}
				tmpChild.ID = child.ID
				tmpChild.Name = child.Name
				tmpChild.ParentID = child.ParentID
				tmpChild.Spacing = 4

				if child.ParentID == 6 && child.ID != 1789 {
					tmpChild.PdfPath = ""
					tmpChild.CoverPath = ""
				} else {
					pdfPath := strings.Split(child.PdfPath, "/")
					coverPath := strings.Split(child.CoverPath, "/")

					tmpChild.PdfPath = pdfPath[len(pdfPath)-1]
					tmpChild.CoverPath = coverPath[len(coverPath)-1]
				}

				if child.ParentID == 6 && child.ID == 1789 && child.CoverPath == "" {
					tmpChild.CoverPath = "un_inggris_smk.jpeg"
					tmpParent.Children = append(tmpParent.Children, tmpChild)
				} else if child.ParentID == 6 && child.ID == 1789 && child.CoverPath != "" {

				} else {
					tmpParent.Children = append(tmpParent.Children, tmpChild)
				}
			}
		}

		results = append(results, tmpParent)
	}

	return results
}

// LoadKisiUNStd load kisi un std
func LoadKisiUNStd() []NationalExamStructure {
	var (
		unStandard     UNStandard
		unStandards    []UNStandard
		resultsWithSub []NationalExamStructure
		results        []NationalExamStructure
		fixResults     []NationalExamStructure
	)

	rows, err := db.Query(unStandardQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch UN standard status")

		return results
	}

	for rows.Next() {
		err = rows.Scan(&unStandard.ID, &unStandard.Name, &unStandard.Sub, &unStandard.PdfPath, &unStandard.CoverPath)
		unStandards = append(unStandards, unStandard)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	for _, item := range unStandards {
		if item.Sub == "-" {
			temp := convertUNStdToNEStruct(item, 91)

			results = append(results, temp)
		} else {
			temp := findKisiUNStds(results, item.ID)

			if temp.ID == 0 {
				temp = convertUNStdToNEStruct(item, 91)
				temp.PdfPath = ""
				temp.CoverPath = ""

				results = append(results, temp)
			}

			tempChild := convertUNStdToNEStruct(item, temp.ID)
			tempChild.Name = item.Sub

			if tempChild.Name == "akuntansi" {
				tempChild.Name = "Akuntansi dan Penjualan"
			} else if tempChild.Name == "pariwisata" {
				tempChild.Name = "Pariwisata, Seni dan Kerajinan, Teknologi Kerumahtanggaan, Pekerjaan Sosial, dan Administrasi Perkantoran"
			} else {
				tempChild.Name = "Teknologi, Kesehatan, dan Pertanian"
			}

			resultsWithSub = append(resultsWithSub, tempChild)
		}
	}

	for _, item := range results {
		item.Spacing = 2
		for _, sub := range resultsWithSub {
			if sub.ParentID == item.ID {
				sub.Spacing = 3
				item.Children = append(item.Children, sub)
			}
		}

		fixResults = append(fixResults, item)
	}

	return fixResults
}

func findKisiUNStds(kisiUNStds []NationalExamStructure, id int) NationalExamStructure {
	for _, item := range kisiUNStds {
		if item.ID == id {
			return item
		}
	}
	return NationalExamStructure{}
}

func convertUNStdToNEStruct(item UNStandard, parentID int) NationalExamStructure {
	result := NationalExamStructure{}

	result.ID = item.ID
	result.Name = item.Name
	result.ParentID = parentID

	pdfPath := strings.Split(item.PdfPath, "/")
	coverPath := strings.Split(item.CoverPath, "/")

	result.CoverPath = coverPath[len(coverPath)-1]
	result.PdfPath = pdfPath[len(pdfPath)-1]

	return result
}

func loadKisiUNUSBNKejuruan(qry string) []NationalExamStructure {
	var (
		flatNationalExamStruct  FlatNationalExamStruct
		flatNationalExamStructs []FlatNationalExamStruct
	)

	rows, err := db.Query(qry)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch structure curriculum status")

		return []NationalExamStructure{}
	}

	for rows.Next() {
		err = rows.Scan(&flatNationalExamStruct.FieldID,
			&flatNationalExamStruct.ProgramID,
			&flatNationalExamStruct.CompetencyID,
			&flatNationalExamStruct.FieldTitle,
			&flatNationalExamStruct.ProgramTitle,
			&flatNationalExamStruct.CompetencyTitle,
			&flatNationalExamStruct.PdfPath,
			&flatNationalExamStruct.JpegCover,
		)

		pdfPath := strings.Split(flatNationalExamStruct.PdfPath, "/")
		coverPath := strings.Split(flatNationalExamStruct.JpegCover, "/")

		flatNationalExamStruct.JpegCover = coverPath[len(coverPath)-1]
		flatNationalExamStruct.PdfPath = pdfPath[len(pdfPath)-1]

		flatNationalExamStructs = append(flatNationalExamStructs, flatNationalExamStruct)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return transformStructureNationalExam(flatNationalExamStructs)
}

// LoadTataTertib fetach tata tertib data
func LoadTataTertib() []NationalExamStructure {
	var (
		root        NationalExamStructure
		unCompBased NationalExamStructure
		usbn        NationalExamStructure
		results     []NationalExamStructure
	)

	unCompBased.ID = 95
	unCompBased.Name = "Ujian Nasional Berbasis Komputer (UNBK) 2019"
	unCompBased.Spacing = 1
	unCompBased.ParentID = 94
	unCompBased.PdfPath = "tata_tertib_peserta_unbk_.pdf"

	usbn.ID = 96
	usbn.Name = "Ujian Sekolah Berstandar Nasional (USBN) 2019"
	usbn.Spacing = 1
	usbn.ParentID = 94
	usbn.PdfPath = "tata_tertib_peserta_usbn.pdf"

	root.ID = 94
	root.Name = "Tata Tertib"
	root.Children = append(root.Children, unCompBased)
	root.Children = append(root.Children, usbn)

	results = append(results, root)

	return results
}

// LoadTimeAllocation fetach alokasi waktu data
func LoadTimeAllocation() []NationalExamStructure {
	var (
		root                    NationalExamStructure
		un                      NationalExamStructure
		usbn                    NationalExamStructure
		muatanNasional          NationalExamStructure
		muatanPeminatanKejuruan NationalExamStructure
		results                 []NationalExamStructure
	)

	un.ID = 98
	un.Name = "Jumlah Butir Soal dan Alokasi Waktu: UN 2019"
	un.Spacing = 1
	un.ParentID = 97
	un.PdfPath = "un_jumlah_soal_smk.pdf"

	muatanNasional.ID = 103
	muatanNasional.Name = "Muatan Nasional"
	muatanNasional.Spacing = 2
	muatanNasional.ParentID = 99
	muatanNasional.PdfPath = "usbn_jumlah_soal_smk_muatan_nasional.pdf"

	muatanPeminatanKejuruan.ID = 104
	muatanPeminatanKejuruan.Name = "Muatan Peminatan Kejuruan"
	muatanPeminatanKejuruan.Spacing = 2
	muatanPeminatanKejuruan.ParentID = 99
	muatanPeminatanKejuruan.PdfPath = "usbn_jumlah_soal_smk_muatan_kejuruan.pdf"

	usbn.ID = 99
	usbn.Name = "Jumlah Butir Soal dan Alokasi Waktu: USBN 2019"
	usbn.Spacing = 1
	usbn.ParentID = 97
	usbn.Children = append(usbn.Children, muatanNasional)
	usbn.Children = append(usbn.Children, muatanPeminatanKejuruan)

	root.ID = 97
	root.Name = "Jumlah Butir Soal dan Alokasi Waktu"
	root.Children = append(root.Children, un)
	root.Children = append(root.Children, usbn)

	results = append(results, root)

	return results
}

// transformStructureNationalExam transform
func transformStructureNationalExam(flatNationalExamStructs []FlatNationalExamStruct) []NationalExamStructure {
	var (
		field        NationalExamStructure
		program      NationalExamStructure
		competency   NationalExamStructure
		fields       []NationalExamStructure
		programs     []NationalExamStructure
		competencies []NationalExamStructure
		result       []NationalExamStructure
	)

	if flatNationalExamStructs != nil {
		for _, item := range flatNationalExamStructs {
			if !nesAlreadyExist(fields, item.FieldID) {
				field.ID = item.FieldID
				//field.ChildID = item.ProgramID
				field.Name = item.FieldTitle
				field.Spacing = 3

				fields = append(fields, field)
			}

			if !nesAlreadyExist(programs, item.ProgramID) {
				program.ID = item.ProgramID
				//program.ChildID = item.CompetencyID
				program.ParentID = item.FieldID
				program.Name = item.ProgramTitle
				program.Spacing = 4

				programs = append(programs, program)
			}

			competency.ID = item.CompetencyID
			competency.ParentID = item.ProgramID
			competency.Name = item.CompetencyTitle
			competency.PdfPath = item.PdfPath
			competency.CoverPath = item.JpegCover
			competency.Spacing = 5

			competencies = append(competencies, competency)
		}

		result = mergeNes(fields, programs, competencies)
	}

	return result
}

func nesAlreadyExist(items []NationalExamStructure, ID int) bool {
	for _, item := range items {
		if item.ID == ID {
			return true
		}
	}

	return false
}

func mergeNes(fields []NationalExamStructure, programs []NationalExamStructure, competencies []NationalExamStructure) []NationalExamStructure {
	var (
		field      NationalExamStructure
		program    NationalExamStructure
		competency NationalExamStructure
		programRes []NationalExamStructure
		result     []NationalExamStructure
	)

	if programs != nil {
		for _, itemProg := range programs {
			program = NationalExamStructure{}
			program.ID = itemProg.ID
			program.Name = itemProg.Name
			program.ParentID = itemProg.ParentID
			program.Spacing = itemProg.Spacing

			for _, itemComp := range competencies {
				if itemComp.ParentID == itemProg.ID {
					competency.ID = itemComp.ID
					competency.Name = itemComp.Name
					competency.ParentID = itemComp.ParentID
					competency.PdfPath = itemComp.PdfPath
					competency.CoverPath = itemComp.CoverPath
					competency.Spacing = itemComp.Spacing

					program.Children = append(program.Children, itemComp)
				}
			}

			programRes = append(programRes, program)
		}
	}

	if fields != nil {
		for _, itemField := range fields {
			field = NationalExamStructure{}
			field.ID = itemField.ID
			field.Name = itemField.Name
			field.ParentID = 100
			field.Spacing = itemField.Spacing

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
