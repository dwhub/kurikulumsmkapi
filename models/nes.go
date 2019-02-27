package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	u "github.com/dwhub/kurikulumsmkapi/utils"
	log "github.com/sirupsen/logrus"
)

// NationalEducationStandard model
type NationalEducationStandard struct {
	ID        int                         `json:"id"`
	ParentID  int                         `json:"parent_id"`
	Name      string                      `json:"name"`
	Spacing   int                         `json:"spacing"`
	HasChild  bool                        `json:"has_child"`
	PdfPath   string                      `json:"pdf_path"`
	CoverPath string                      `json:"cover_path"`
	Children  []NationalEducationStandard `json:"children"`
}

// FlatIsiJurusan model
type FlatIsiJurusan struct {
	BidangID                  int            `json:"bidang_id"`
	ProgramID                 int            `json:"program_id"`
	TitleBidang               string         `json:"title_bidang"`
	TitleProgram              string         `json:"title_program"`
	StandarUmumFile           string         `json:"standar_umum_file"`
	StandarKejuruanFile       sql.NullString `json:"standar_kejuruan_file"`
	StandarKejuruanBidangFile sql.NullString `json:"standar_kejuruan_bidang_file"`
}

var standardRoot = `SELECT standar_id, CONCAT(standar_urutan, '. ', standar_judul) as title, standar_has_sub_1, standar_file
					FROM tbl_standar 
					ORDER BY standar_urutan`

var standarIsiQuery = `SELECT DISTINCT a.id_bidang, b.id_program, 
							CONCAT(a.urutan_bidang, '. ', a.bidang_keahlian) as title_bidang, 
							CONCAT(a.urutan_bidang, '.', b.urutan_program, '. ', b.program_keahlian) as title_program,
							c.standar_umum_file, d.standar_kejuruan_file, e.standar_kejuruan_file as tbl_standar_kejuruan_bidang
							FROM tbl_bidang_keahlian a 
							LEFT JOIN tbl_program_keahlian b on a.id_bidang = b.id_bidang
							LEFT JOIN tbl_standar_umum c ON a.id_bidang = c.id_bidang
							LEFT JOIN tbl_standar_kejuruan d ON b.id_program = d.id_program
							LEFT JOIN tbl_standar_kejuruan_bidang e ON a.id_bidang = e.id_bidang
							ORDER BY a.id_bidang, b.id_program`

var standarKeahlianQuery = `SELECT a.standar_sub_1_id, a.standar_id_parent,
							CONCAT(a.standar_sub_1_urutan, '. ', a.standar_sub_1_judul) as title_sub, 
							a.standar_sub_1_file,
							a.standar_sub_1_is_jurusan,
							b.standar_sub_1_id_parent,
							b.standar_sub_2_judul,
							b.standar_sub_2_file
							FROM tbl_standar_sub_1 a 
							LEFT JOIN tbl_standar_sub_2 b ON a.standar_sub_1_id = b.standar_sub_1_id_parent
							ORDER BY a.standar_sub_1_id`

var standarKeahlianJurusanQuery = `SELECT a.id_bidang, b.id_program, c.id_kompetensi, 
									CONCAT(a.urutan_bidang, '. ', a.bidang_keahlian) as title_bidang, 
									CONCAT(a.urutan_bidang, '.', b.urutan_program, '. ', b.program_keahlian) as title_program, 
									CONCAT(a.urutan_bidang, '.', b.urutan_program, '.', c.urutan_kompetensi, '. ', c.kompetensi_keahlian) as title_kompetensi,
									d.standar_keahlian_file
									FROM tbl_bidang_keahlian a 
									LEFT JOIN tbl_program_keahlian b on a.id_bidang = b.id_bidang
									LEFT JOIN tbl_kompetensi_keahlian c on b.id_program = c.id_program
									LEFT JOIN tbl_standar_keahlian d on c.id_kompetensi = d.id_kompetensi`

// GetNationalEducationStandards fetch all national education standard data without paging
func GetNationalEducationStandards() map[string]interface{} {
	var (
		root []NationalEducationStandard
	)

	root = LoadRoot()

	return u.Message(http.StatusOK, root)
}

// LoadRoot ...
func LoadRoot() []NationalEducationStandard {
	var (
		nationalEducationStandard NationalEducationStandard
		root                      []NationalEducationStandard
		result                    []NationalEducationStandard
	)

	rows, err := db.Query(standardRoot)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch national education field status")
	}

	for rows.Next() {
		err = rows.Scan(&nationalEducationStandard.ID, &nationalEducationStandard.Name,
			&nationalEducationStandard.HasChild, &nationalEducationStandard.PdfPath)
		root = append(root, nationalEducationStandard)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	isiStandar := ProcessStandarIsi()
	for _, item := range root {
		temp := createCopyNES(item)

		if item.HasChild && item.ID == 2 {
			for _, bidang := range isiStandar {
				tempChild := createCopyNES(bidang)

				temp.Children = append(temp.Children, tempChild)
			}
		}

		result = append(result, temp)
	}

	return result
}

// ProcessStandarIsi method
func ProcessStandarIsi() []NationalEducationStandard {
	var (
		isiJurusan  []FlatIsiJurusan
		muatan      []NationalEducationStandard
		results     []NationalEducationStandard
		tempResults []NationalEducationStandard
	)

	isiJurusan = LoadIsiJurusan()
	muatan = createMuatan()

	for _, item := range isiJurusan {
		temp := NationalEducationStandard{}

		for _, itemMuatan := range muatan {
			itemMuatan.ParentID = item.BidangID
			if itemMuatan.ID == 90 {
				itemMuatan.PdfPath = findMuatanUmumPdf(item.BidangID, isiJurusan)
			}

			if itemMuatan.ID == 91 {
				muatanChild := createMuatanKejuruan(item.BidangID, isiJurusan)

				for _, itemMuatanChild := range muatanChild {
					itemMuatan.Children = append(itemMuatan.Children, itemMuatanChild)
				}

				if len(muatanChild) == 0 {
					itemMuatan.HasChild = false
					itemMuatan.PdfPath = findMuatanKejuruanBidangPdf(item.BidangID, isiJurusan)[0]
				}
			}
			tempMuatan := createCopyNES(itemMuatan)

			temp.Children = append(temp.Children, tempMuatan)
		}
		temp.ID = item.BidangID
		temp.Name = item.TitleBidang
		temp.Spacing = 1
		temp.ParentID = 2

		tempResults = append(tempResults, temp)
	}

	for _, item := range tempResults {
		exist := false
		for _, bidang := range results {
			if bidang.ID == item.ID {
				exist = true
			}
		}

		if !exist {
			results = append(results, item)
		}
	}
	return results
}

func createMuatanKejuruan(ID int, isiJurusan []FlatIsiJurusan) []NationalEducationStandard {
	var (
		temp   []NationalEducationStandard
		result []NationalEducationStandard
	)

	for _, item := range isiJurusan {
		pdf := findMuatanKejuruanPdf(item.ProgramID, isiJurusan)

		if item.BidangID == ID && pdf[0] != "" {
			tempProgram := NationalEducationStandard{}

			if len(pdf) > 1 {
				tempProgramTahun := createMuatanKejuruanTahun()

				for _, itemTahun := range tempProgramTahun {
					itemTahun.ParentID = item.ProgramID
					if itemTahun.ID == 92 {
						itemTahun.PdfPath = pdf[0]
					} else {
						itemTahun.PdfPath = pdf[1]
					}

					tempItemTahun := createCopyNES(itemTahun)

					tempProgram.HasChild = true
					tempProgram.Children = append(tempProgram.Children, tempItemTahun)
				}
			} else {
				tempProgram.PdfPath = pdf[0]
			}

			tempProgram.ID = item.ProgramID
			tempProgram.Name = item.TitleProgram
			tempProgram.Spacing = 3
			tempProgram.ParentID = 91

			temp = append(temp, tempProgram)
		}
	}

	for _, item := range temp {
		exist := false
		for _, program := range result {
			if item.ID == program.ID {
				exist = true
			}
		}

		if !exist {
			tempProgram := createCopyNES(item)
			result = append(result, tempProgram)
		}
	}
	return result
}

func createMuatan() []NationalEducationStandard {
	var result []NationalEducationStandard

	muatanUmum := NationalEducationStandard{}
	muatanUmum.ID = 90
	muatanUmum.HasChild = false
	muatanUmum.Name = "A. Kompetensi Muatan Umum"
	muatanUmum.Spacing = 2

	muatanKejuruan := NationalEducationStandard{}
	muatanKejuruan.ID = 91
	muatanKejuruan.HasChild = true
	muatanKejuruan.Name = "B. Kompetensi Muatan Kejuruan"
	muatanKejuruan.Spacing = 2

	result = append(result, muatanUmum)
	result = append(result, muatanKejuruan)

	return result
}

func createMuatanKejuruanTahun() []NationalEducationStandard {
	var result []NationalEducationStandard

	tigaTahun := NationalEducationStandard{}
	tigaTahun.ID = 92
	tigaTahun.HasChild = false
	tigaTahun.Name = "1) Program 3 (tiga) Tahun"
	tigaTahun.Spacing = 4

	empatTahun := NationalEducationStandard{}
	empatTahun.ID = 93
	empatTahun.HasChild = false
	empatTahun.Name = "2) Program 4 (empat) Tahun"
	empatTahun.Spacing = 4

	result = append(result, tigaTahun)
	result = append(result, empatTahun)

	return result
}

// LoadIsiJurusan method
func LoadIsiJurusan() []FlatIsiJurusan {
	var (
		jurusan    FlatIsiJurusan
		isiJurusan []FlatIsiJurusan
	)

	rows, err := db.Query(standarIsiQuery)

	if err != nil {
		log.WithFields(log.Fields{
			"status": "Failed",
			"error":  err,
		}).Info("Fetch national education field status")
	}

	for rows.Next() {
		err = rows.Scan(&jurusan.BidangID, &jurusan.ProgramID, &jurusan.TitleBidang,
			&jurusan.TitleProgram, &jurusan.StandarUmumFile,
			&jurusan.StandarKejuruanFile, &jurusan.StandarKejuruanBidangFile)
		isiJurusan = append(isiJurusan, jurusan)
		if err != nil {
			fmt.Print(err.Error())
		}
	}

	return isiJurusan
}

func createCopyNES(item NationalEducationStandard) NationalEducationStandard {
	var result NationalEducationStandard

	result.ID = item.ID
	result.Name = item.Name
	result.ParentID = item.ParentID
	result.PdfPath = item.PdfPath
	result.Spacing = item.Spacing
	result.CoverPath = item.CoverPath
	result.HasChild = item.HasChild

	if len(item.Children) > 0 {
		for _, child := range item.Children {
			temp := createCopyNES(child)

			result.Children = append(result.Children, temp)
		}
	}
	return result
}

func findMuatanUmumPdf(ID int, isiJurusan []FlatIsiJurusan) string {
	var result string

	for _, item := range isiJurusan {
		if item.BidangID == ID {
			pdfPath := strings.Split(item.StandarUmumFile, "/")

			result = pdfPath[len(pdfPath)-1]
			break
		}
	}
	return result
}

func findMuatanKejuruanPdf(ID int, isiJurusan []FlatIsiJurusan) []string {
	var result []string

	for _, item := range isiJurusan {
		if item.ProgramID == ID {
			pdfPath := strings.Split(item.StandarKejuruanFile.String, "/")

			result = append(result, pdfPath[len(pdfPath)-1])
		}
	}
	return result
}

func findMuatanKejuruanBidangPdf(ID int, isiJurusan []FlatIsiJurusan) []string {
	var result []string

	for _, item := range isiJurusan {
		if item.BidangID == ID {
			pdfPath := strings.Split(item.StandarKejuruanBidangFile.String, "/")

			result = append(result, pdfPath[len(pdfPath)-1])
		}
	}
	return result
}
