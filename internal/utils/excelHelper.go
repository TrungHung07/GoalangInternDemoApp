package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelClassReport struct {
	STT             int
	Class           string
	Students        []string
	Teachers        []string
	StudentQuantity int
	TeacherQuantity int
}

type ExcelHelper struct{}

func NewExcelHelper() *ExcelHelper {
	return &ExcelHelper{}
}

func (e *ExcelHelper) ExportClassTableReport(title string, data []ExcelClassReport) ([]byte, error) {
	f := excelize.NewFile()
	defer f.Close()
	sheet := "Report"
	f.SetSheetName("Sheet1", sheet)

	row := 1
	// Title
	titleCell := fmt.Sprintf("A%d", row)
	f.SetCellValue(sheet, titleCell, title)
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 16},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.MergeCell(sheet, "A1", "F1")
	f.SetCellStyle(sheet, "A1", "F1", titleStyle)
	row += 1

	// Headers
	headers := []string{"STT", "Class", "Student", "Teacher", "Student Quantity", "Teacher Quantity"}
	for i, h := range headers {
		f.SetCellValue(sheet, fmt.Sprintf("%s%d", e.col(i), row), h)
	}
	topHeaderRow := row
	row++

	// Table content
	for _, entry := range data {
		maxRows := max(len(entry.Students), len(entry.Teachers))
		startRow := row
		endRow := row + maxRows - 1

		// Merge STT, Class, Student Quantity, Teacher Quantity
		f.MergeCell(sheet, fmt.Sprintf("A%d", startRow), fmt.Sprintf("A%d", endRow))
		f.SetCellValue(sheet, fmt.Sprintf("A%d", startRow), entry.STT)

		f.MergeCell(sheet, fmt.Sprintf("B%d", startRow), fmt.Sprintf("B%d", endRow))
		f.SetCellValue(sheet, fmt.Sprintf("B%d", startRow), entry.Class)

		f.MergeCell(sheet, fmt.Sprintf("E%d", startRow), fmt.Sprintf("E%d", endRow))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", startRow), entry.StudentQuantity)

		f.MergeCell(sheet, fmt.Sprintf("F%d", startRow), fmt.Sprintf("F%d", endRow))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", startRow), entry.TeacherQuantity)

		// Students and Teachers
		for i := 0; i < maxRows; i++ {
			if i < len(entry.Students) {
				f.SetCellValue(sheet, fmt.Sprintf("C%d", row), entry.Students[i])
			}
			if i < len(entry.Teachers) {
				f.SetCellValue(sheet, fmt.Sprintf("D%d", row), entry.Teachers[i])
			}
			row++
		}
	}

	// Apply border style
	borderStyle, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})

	f.SetCellStyle(sheet, fmt.Sprintf("A%d", topHeaderRow), fmt.Sprintf("F%d", row-1), borderStyle)
	f.SetColWidth(sheet, "A", "F", 20)

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (e *ExcelHelper) col(idx int) string {
	col := ""
	for idx >= 0 {
		col = string(rune('A'+idx%26)) + col
		idx = idx/26 - 1
	}
	return col
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
