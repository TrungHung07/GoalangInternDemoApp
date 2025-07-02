package utils

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelSection struct {
	Title    string              `json:"title"`
	Headers  []string            `json:"headers"`
	Data     []map[string]string `json:"data"`
	StartRow int                 `json:"start_row"`
}

type ExcelReportData struct {
	Title     string                 `json:"title"`
	BasicInfo map[string]interface{} `json:"basic_info"`
	Sections  []ExcelSection         `json:"sections"`
	SheetName string                 `json:"sheet_name"`
}

type ExcelHelper struct{}

func NewExcelHelper() *ExcelHelper {
	return &ExcelHelper{}
}

func (e *ExcelHelper) ExportReport(data *ExcelReportData) ([]byte, error) {
	file := excelize.NewFile()
	defer file.Close()

	sheetName := data.SheetName
	if sheetName == "" {
		sheetName = "Report"
	}
	file.SetSheetName("Sheet1", sheetName)

	row := 1
	if data.Title != "" {
		row = e.setTitle(file, sheetName, data.Title, row)
		row += 2
	}

	if len(data.BasicInfo) > 0 {
		row = e.setBasicInfo(file, sheetName, data.BasicInfo, row)
		row += 2
	}

	for _, section := range data.Sections {
		row = e.setSection(file, sheetName, section, row)
		row += 2
	}

	e.applyStyles(file, sheetName)

	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("failed to export Excel: %v", err)
	}
	return buf.Bytes(), nil
}

func (e *ExcelHelper) setTitle(file *excelize.File, sheet, title string, row int) int {
	cell := fmt.Sprintf("A%d", row)
	file.SetCellValue(sheet, cell, title)
	startCell := fmt.Sprintf("A%d", row)
	endCell := fmt.Sprintf("D%d", row)
	file.MergeCell(sheet, startCell, endCell)
	return row + 1
}

func (e *ExcelHelper) setBasicInfo(file *excelize.File, sheet string, info map[string]interface{}, row int) int {
	for k, v := range info {
		file.SetCellValue(sheet, fmt.Sprintf("A%d", row), k+":")
		file.SetCellValue(sheet, fmt.Sprintf("B%d", row), v)
		row++
	}
	return row
}

func (e *ExcelHelper) setSection(file *excelize.File, sheet string, section ExcelSection, row int) int {
	if section.Title != "" {
		file.SetCellValue(sheet, fmt.Sprintf("A%d", row), section.Title)
		row += 2
	}

	for i, header := range section.Headers {
		file.SetCellValue(sheet, fmt.Sprintf("%s%d", e.col(i), row), header)
	}
	row++

	for _, dataRow := range section.Data {
		for i, header := range section.Headers {
			if val, ok := dataRow[header]; ok {
				file.SetCellValue(sheet, fmt.Sprintf("%s%d", e.col(i), row), val)
			}
		}
		row++
	}
	return row
}

func (e *ExcelHelper) applyStyles(file *excelize.File, sheet string) {
	titleStyle, _ := file.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 16, Family: "Arial"},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})

	headerStyle, _ := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 12, Family: "Arial"},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#E6E6FA"}, Pattern: 1},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})

	file.SetColWidth(sheet, "A", "D", 20)
	_ = titleStyle
	_ = headerStyle
}

func (e *ExcelHelper) col(idx int) string {
	col := ""
	for idx >= 0 {
		col = string(rune('A'+idx%26)) + col
		idx = idx/26 - 1
	}
	return col
}
