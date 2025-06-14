package tabeludo_test

import (
	"testing"

	"github.com/SamuelLFA/tabeludo"
	"github.com/stretchr/testify/assert"
)

const _testfilepath = "test_files/1.xlsx"

func TestNewReader(t *testing.T) {
	reader, err := tabeludo.NewReader(_testfilepath)
	assert.Nil(t, err)

	assert.NotNil(t, reader)
	reader.Close()
}

func TestClose(t *testing.T) {
	reader, err := tabeludo.NewReader(_testfilepath)
	assert.Nil(t, err)

	err = reader.Close()
	assert.Nil(t, err)
	err = reader.Close()
	assert.Equal(t, "close test_files/1.xlsx: file already closed", err.Error())
}
