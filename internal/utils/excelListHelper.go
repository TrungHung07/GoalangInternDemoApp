package utils

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuri/excelize/v2"
)

type ExcelListHelper struct{}

func NewExcelListHelper() *ExcelListHelper {
	return &ExcelListHelper{}
}

type ExcelClassData struct {
	Class            ClassInfo
	Students         []StudentInfo
	Teachers         []TeacherInfo
	StudentsQuantity int
	TeachersQuantity int
}

type ClassInfo struct {
	ID    int
	Name  string
	Grade int
}

type StudentInfo struct {
	Name string
}

type TeacherInfo struct {
	Name  string
	Email string
	Age   int
}

func (e *ExcelListHelper) ExportClassListExcel(data []ExcelClassData) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName("Sheet1", "Báo Cáo")
	sheet = "Báo Cáo"

	// Title
	f.SetCellValue(sheet, "A1", "BÁO CÁO ABC")
	f.MergeCell(sheet, "A1", "H1")
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 16},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.SetCellStyle(sheet, "A1", "H1", titleStyle)

	// Header
	headers := []string{"STT", "Class", "Student", "Teacher Name", "Teacher Email", "Teacher Age", "Student Quantity", "Teacher Quantity"}
	for i, h := range headers {
		col := setCol(i)
		f.SetCellValue(sheet, fmt.Sprintf("%s2", col), h)
	}
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Border: []excelize.Border{
			{Type: "left", Style: 1},
			{Type: "top", Style: 1},
			{Type: "bottom", Style: 1},
			{Type: "right", Style: 1},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetCellStyle(sheet, "A2", "H2", headerStyle)

	row := 3
	for i, classData := range data {
		maxLen := maxValue(len(classData.Students), len(classData.Teachers))
		startRow := row
		endRow := row + maxLen - 1

		// STT
		f.SetCellValue(sheet, fmt.Sprintf("A%d", startRow), i+1)
		if maxLen > 1 {
			f.MergeCell(sheet, fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", endRow))
		}

		// Class
		f.SetCellValue(sheet, fmt.Sprintf("B%d", startRow), classData.Class.Name)
		if maxLen > 1 {
			f.MergeCell(sheet, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", endRow))
		}

		// Students
		for sIdx := 0; sIdx < maxLen; sIdx++ {
			if sIdx < len(classData.Students) {
				f.SetCellValue(sheet, fmt.Sprintf("C%d", row+sIdx), classData.Students[sIdx].Name)
			}
		}

		// Teachers
		for tIdx := 0; tIdx < maxLen; tIdx++ {
			if tIdx < len(classData.Teachers) {
				t := classData.Teachers[tIdx]
				f.SetCellValue(sheet, fmt.Sprintf("D%d", row+tIdx), t.Name)
				f.SetCellValue(sheet, fmt.Sprintf("E%d", row+tIdx), t.Email)
				f.SetCellValue(sheet, fmt.Sprintf("F%d", row+tIdx), t.Age)
			}
		}

		// Quantities
		f.SetCellValue(sheet, fmt.Sprintf("G%d", startRow), classData.StudentsQuantity)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", startRow), classData.TeachersQuantity)
		if maxLen > 1 {
			f.MergeCell(sheet, fmt.Sprintf("G%d", startRow), fmt.Sprintf("G%d", endRow))
			f.MergeCell(sheet, fmt.Sprintf("H%d", startRow), fmt.Sprintf("H%d", endRow))
		}

		row += maxLen
	}

	f.SetColWidth(sheet, "A", "H", 20)

	dataEndRow := row - 1
	log.Infof("end row : %d", dataEndRow)
	if dataEndRow < 2 {
		dataEndRow = 2
	}
	// cellRange := fmt.Sprintf("A1:H%d", dataEndRow)
	startCell := "A1"
	endCell := fmt.Sprintf("H%d", dataEndRow)
	combinedStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	f.SetCellStyle(sheet, startCell, endCell, combinedStyle)

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// setCol converts 0 => A, 1 => B, ..., 25 => Z, 26 => AA, ...
func setCol(idx int) string {
	col := ""
	for idx >= 0 {
		col = string(rune('A'+idx%26)) + col
		idx = idx/26 - 1
	}
	return col
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
