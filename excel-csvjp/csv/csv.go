package csv

import (
	"encoding/csv"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// NewReader 与えられたio.Readerを元に新しいcsvリーダーを返す
func NewReader(r io.Reader) *csv.Reader {
	return csv.NewReader(transform.NewReader(r, japanese.ShiftJIS.NewDecoder()))
}

// NewWriter 与えられたio.Writerを元に新しいcsvライターを返す。UseCRLFはデフォルトでtrueが設定される
func NewWriter(w io.Writer) *csv.Writer {
	writer := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
	writer.UseCRLF = true
	return writer
}
