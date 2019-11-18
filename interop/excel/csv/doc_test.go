package csv_test

import (
	"fmt"
	"os"

	"github.com/topgate/goutils/interop/excel/csv"
)

func Example_writer() {
	file, err := os.Create("some.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewUTF8WithBOMWriter(file)
	defer writer.Flush()
	writer.WriteAll([][]string{
		[]string{"header1", "header2"},
		[]string{"val1", "val2"},
	})
}

func Example_reader() {
	file, err := os.Open("some_sjis_utf8withbom.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", data)
}

func ExampleNewUTF8WithBOMReader() {
	file, err := os.Open("some.csv")
	if err != nil {
		panic(err)
	}
	csv.NewUTF8WithBOMReader(file)
	// Output: BOM付きUTF8が読み取り可能な *csv.Reader
}

func ExampleNewSJISReader() {
	file, err := os.Open("some.csv")
	if err != nil {
		panic(err)
	}
	csv.NewSJISReader(file)
	// Output: SJISが読み取り可能な *csv.Reader
}

func ExampleNewReader() {
	file, err := os.Open("some.csv")
	if err != nil {
		panic(err)
	}
	csv.NewReader(file)
	// Output: SJIS,BOM付きUTF8両方が読み取り可能な *csv.Reader
}

func ExampleNewSJISWriter() {
	file, err := os.Open("some.csv")
	if err != nil {
		panic(err)
	}
	csv.NewSJISWriter(file)
	// Output: SJISで書き込み可能な *csv.Writer

}

func ExampleNewUTF8WithBOMWriter() {
	file, err := os.Open("some.csv")
	if err != nil {
		panic(err)
	}
	csv.NewUTF8WithBOMWriter(file)
	// Output: BOM付きUTF8で書き込み可能な *csv.Writer

}
