package gocsv

import (
	"io"

	"github.com/gocarina/gocsv"
	"github.com/topgate/goutils/interop/excel/csv"
)

// Unmarshal parses the CSV from the reader in the interface.
func Unmarshal(r io.Reader, out interface{}) error {
	return gocsv.UnmarshalCSV(csv.NewReader(r), out)
}

// Marshal returns  CSV in writer from the interface.
func Marshal(in interface{}, w io.Writer) error {
	return gocsv.MarshalCSV(in, gocsv.NewSafeCSVWriter(csv.NewWriter(w)))
}
