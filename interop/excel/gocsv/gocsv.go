// Package gocsv Excelと相互運用できるCSVを扱うgocsvと関数名の互換性があるパッケージ
package gocsv

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/topgate/goutils/interop/excel/csv"
)

// Unmarshal gocsvを使用して、rから読み取られたデータをoutにバインドする
func Unmarshal(r io.Reader, out interface{}) error {
	return gocsv.UnmarshalCSV(csv.NewReader(r), out)
}

// Marshal returns  gocsvを使用して、inをwで与えられたWriterに書き込む
func Marshal(in interface{}, w io.Writer) error {
	return gocsv.MarshalCSV(in, gocsv.NewSafeCSVWriter(csv.NewUTF8WithBOMWriter(w)))
}
