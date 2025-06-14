package tabeludo

import (
	"archive/zip"
	"fmt"
)

type Reader struct {
	filename  string
	zipReader *zip.ReadCloser
}

func NewReader(filename string) (*Reader, error) {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, fmt.Errorf("opening XLSX file '%s': %v", filename, err)
	}

	reader := &Reader{
		filename:  filename,
		zipReader: zipReader,
	}

	return reader, nil
}

func (r *Reader) Close() error {
	if r.zipReader != nil {
		return r.zipReader.Close()
	}

	return nil
}
