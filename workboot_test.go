package tabeludo_test

import (
	"testing"

	"github.com/SamuelLFA/tabeludo"
	"github.com/stretchr/testify/assert"
)

func TestWorkbookReadWorkbook(t *testing.T) {
	r, err := tabeludo.NewReader(_testfilepath)
	assert.Nil(t, err)

	wb, err := r.ReadWorkbook()
	assert.Nil(t, err)

	assert.Equal(t, 1, len(wb.Sheets))
	assert.Equal(t, "Sheet1", wb.Sheets[0].Name)
	assert.Equal(t, "1", wb.Sheets[0].SheetID)

	r.Close()
}
