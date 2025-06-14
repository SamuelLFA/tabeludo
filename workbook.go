package tabeludo

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

type Workbook struct {
	Sheets []Sheet
}

type Sheet struct {
	Name    string
	SheetID string
}

// XML
type xlsxWorkbook struct {
	XMLName xml.Name   `xml:"workbook"`
	Sheets  xlsxSheets `xml:"sheets"`
}

type xlsxSheets struct {
	Sheet []xlsxSheet `xml:"sheet"`
}

type xlsxSheet struct {
	Name    string `xml:"name,attr"`
	SheetID string `xml:"sheetId,attr"`
}

func (r *Reader) ReadWorkbook() (*Workbook, error) {
	workbook, err := r.readWorkbookStructure()
	if err != nil {
		return nil, err
	}

	wb := &Workbook{}
	for _, sheetInfo := range workbook.Sheets.Sheet {
		sheet := Sheet{
			Name:    sheetInfo.Name,
			SheetID: sheetInfo.SheetID,
		}
		wb.Sheets = append(wb.Sheets, sheet)
	}

	return wb, nil
}

func (r *Reader) readWorkbookStructure() (*xlsxWorkbook, error) {
	file := r.findFile("xl/workbook.xml")
	if file == nil {
		return nil, fmt.Errorf("workbook.xml file not found")
	}

	rc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("opening workbook.xml: %v", err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("reading workbook.xml: %v", err)
	}

	var workbook xlsxWorkbook
	err = xml.Unmarshal(data, &workbook)
	if err != nil {
		return nil, fmt.Errorf("parsing workbook.xml: %v", err)
	}

	return &workbook, nil
}

func (r *Reader) findFile(name string) *zip.File {
	for _, file := range r.zipReader.File {
		if file.Name == name {
			return file
		}
	}
	return nil
}
