// Package csv エクセルとCSVを相互運用するためのパッケージ
package csv

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"

	"bytes"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// UTF8BOM UTF8のBOM
var UTF8BOM = [3]byte{0xEF, 0xBB, 0xBF}

// NewUTF8WithBOMReader BOM付きUTF8のリーダーを新しく作る
func NewUTF8WithBOMReader(r io.Reader) *csv.Reader {
	return csv.NewReader(&utf8WIthBOMByteReader{
		readBOM: false,
		reader:  r,
	})
}

// NewSJISReader 与えられたio.Readerを元に新しいリーダーを返す
func NewSJISReader(r io.Reader) *csv.Reader {
	return csv.NewReader(transform.NewReader(r, japanese.ShiftJIS.NewDecoder()))
}

// NewReader 与えられたio.Readerを元に新しいcsvリーダーを返す
func NewReader(r io.Reader) *csv.Reader {
	return csv.NewReader(&hybridByteReader{
		readBOM: false,
		reader:  r,
	})
}

// NewSJISWriter 与えられたio.Writerを元に新しいSJISのcsvライターを返す。
//
// UseCRLFはデフォルトでtrueが設定される
//
// panicが発生するため*bufio.Writerをwに与えてはならない。
// *bufio.Writerを与えるとpanicになるのは、返されるWriterのFlushを呼び出しても、与えた*bufio.WriterのFlushが呼び出されないため。
// これはencoding/csvのWriterの設計に依存する制約である。
func NewSJISWriter(w io.Writer) *csv.Writer {
	if _, ok := w.(*bufio.Writer); ok {
		panic("Can't use *bufio.Writer")
	}
	writer := csv.NewWriter(transform.NewWriter(w, japanese.ShiftJIS.NewEncoder()))
	writer.UseCRLF = true
	return writer
}

// NewUTF8WithBOMWriter 与えられたio.Writerを元に新しいBOM付きUTF8のcsvライターを返す。
//
// UseCRLFはデフォルトでtrueが設定される。
//
// panicが発生するため*bufio.Writerをwに与えてはならない。
// *bufio.Writerを与えるとpanicになるのは、返されるWriterのFlushを呼び出しても、与えた*bufio.WriterのFlushが呼び出されないため。
// これはencoding/csvのWriterの設計に依存する制約である。
func NewUTF8WithBOMWriter(w io.Writer) *csv.Writer {
	if _, ok := w.(*bufio.Writer); ok {
		panic("Can't use *bufio.Writer")
	}
	writer := csv.NewWriter(&utf8WithBOMByteWriter{
		wroteBOM: false,
		writer:   w,
	})
	writer.UseCRLF = true
	return writer
}

type utf8WithBOMByteWriter struct {
	wroteBOM bool
	writer   io.Writer
}

func (u *utf8WithBOMByteWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if !u.wroteBOM {
		n, err := u.writer.Write(UTF8BOM[:])
		u.wroteBOM = true
		if err != nil {
			return n, err
		}
	}
	return u.writer.Write(p)
}

type utf8WIthBOMByteReader struct {
	readBOM bool
	reader  io.Reader
}

func (u *utf8WIthBOMByteReader) Read(p []byte) (int, error) {
	if !u.readBOM {
		bom := [len(UTF8BOM)]byte{}
		n, err := u.reader.Read(bom[:])
		u.readBOM = true
		if err != nil {
			return n, err
		}
		if !bytes.Equal(UTF8BOM[:], bom[:]) {
			return n, errors.New("Invalid UTF8 BOM")

		}
	}
	return u.reader.Read(p)
}

type hybridByteReader struct {
	readBOM bool
	reader  io.Reader
}

func (h *hybridByteReader) Read(p []byte) (int, error) {
	if !h.readBOM {
		bom := [len(UTF8BOM)]byte{}
		n, err := h.reader.Read(bom[:])
		h.readBOM = true
		if err != nil {
			return n, err
		}
		if n < len(UTF8BOM) || !bytes.Equal(UTF8BOM[:], bom[:]) {
			h.reader = transform.NewReader(io.MultiReader(bytes.NewReader(bom[:n]), h.reader), japanese.ShiftJIS.NewDecoder())
		}
	}
	return h.reader.Read(p)
}
