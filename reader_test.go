package tabeludo_test

import (
	"testing"

	"github.com/SamuelLFA/tabeludo"
	"github.com/stretchr/testify/assert"
)

const _validfile = "test_files/valid_file.xlsx"
const _hugeworksheetname = "test_files/huge_worksheet_name.xlsx"

func TestNewReader_ValidFile(t *testing.T) {
	r, err := tabeludo.NewReader(_validfile)
	assert.Nil(t, err)

	assert.NotNil(t, r)
	r.Close()
}

func TestNewReader_NotFound(t *testing.T) {
	_, err := tabeludo.NewReader("")
	assert.Equal(t, "opening XLSX file '': open : no such file or directory", err.Error())
}

func TestClose(t *testing.T) {
	r, err := tabeludo.NewReader(_validfile)
	assert.Nil(t, err)

	err = r.Close()
	assert.Nil(t, err)
	err = r.Close()
	assert.Equal(t, "close test_files/1.xlsx: file already closed", err.Error())
}
