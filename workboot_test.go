package tabeludo_test

import (
	"testing"

	"github.com/SamuelLFA/tabeludo"
	"github.com/stretchr/testify/assert"
)

func TestWorkbookReadWorkbook_ValidFile(t *testing.T) {
	// Valid file
	r, err := tabeludo.NewReader(_validfile)
	assert.Nil(t, err)

	wb, err := r.ReadWorkbook()
	assert.Nil(t, err)

	assert.Equal(t, 1, len(wb.Sheets))
	assert.Equal(t, "Sheet1", wb.Sheets[0].Name)
	assert.Equal(t, "1", wb.Sheets[0].SheetID)

	r.Close()
}

func TestWorkbookReadWorkbook_HugeWorksheetName(t *testing.T) {
	// Huge Worksheet Name
	r, err := tabeludo.NewReader(_hugeworksheetname)
	assert.Nil(t, err)

	wb, err := r.ReadWorkbook()
	assert.Nil(t, err)

	assert.Equal(t, 1, len(wb.Sheets))
	assert.Equal(t, "ThisIsAGiantNameForAWorksheetName", wb.Sheets[0].Name)
	assert.Equal(t, "1", wb.Sheets[0].SheetID)

	r.Close()
}
