package csv

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
)

func TestNewReaderAsSJIS(t *testing.T) {
	var tests = []struct {
		name     string
		expected [][]string
		given    string
	}{
		{
			"最終行改行あり",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
		},
		{
			"最終行改行なし",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3",
		},
		{
			"empty",
			nil,
			"",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			encoder := japanese.ShiftJIS.NewEncoder()
			given, err := encoder.Bytes([]byte(tt.given))
			if !assertions.NoError(err) {
				return
			}
			reader := NewReader(bytes.NewReader(given))
			if !assertions.NoError(err) {
				return
			}
			actual, err := reader.ReadAll()
			if !assertions.NoError(err) {
				return
			}
			assertions.EqualValues(tt.expected, actual)
		})
	}
}

func TestNewSJISReader(t *testing.T) {
	var tests = []struct {
		name     string
		expected [][]string
		given    string
	}{
		{
			"最終行改行あり",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
		},
		{
			"最終行改行なし",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3",
		},
		{
			"empty",
			nil,
			"",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			encoder := japanese.ShiftJIS.NewEncoder()
			given, err := encoder.Bytes([]byte(tt.given))
			if !assertions.NoError(err) {
				return
			}
			reader := NewSJISReader(bytes.NewReader(given))
			if !assertions.NoError(err) {
				return
			}
			actual, err := reader.ReadAll()
			if !assertions.NoError(err) {
				return
			}
			assertions.EqualValues(tt.expected, actual)
		})
	}
}
func TestNewReaderASUTF8WithBOM(t *testing.T) {

	var tests = []struct {
		name     string
		expected [][]string
		given    string
	}{
		{
			"最終行改行あり",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
		},
		{
			"最終行改行なし",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3",
		},
		{
			"empty",
			nil,
			"",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			given := append(UTF8BOM[:], []byte(tt.given)...)
			reader := NewReader(bytes.NewReader(given))
			actual, err := reader.ReadAll()
			if !assertions.NoError(err) {
				return
			}
			assertions.EqualValues(tt.expected, actual)
		})
	}
}

func TestNewUTF8WithBOMReader(t *testing.T) {

	var tests = []struct {
		name     string
		expected [][]string
		given    string
	}{
		{
			"最終行改行あり",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
		},
		{
			"最終行改行なし",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3",
		},
		{
			"empty",
			nil,
			"",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			given := append(UTF8BOM[:], []byte(tt.given)...)
			reader := NewUTF8WithBOMReader(bytes.NewReader(given))
			actual, err := reader.ReadAll()
			if !assertions.NoError(err) {
				return
			}
			assertions.EqualValues(tt.expected, actual)
		})
	}
}
func TestNewSJISWriter(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    [][]string
	}{
		{
			"正常系",
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
		},
		{
			"empty",
			"",
			nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			iw := bytes.Buffer{}
			writer := NewSJISWriter(&iw)
			err := writer.WriteAll(tt.given)
			if !assertions.NoError(err) {
				return
			}
			var expected []byte
			if len(tt.expected) > 0 {
				encoder := japanese.ShiftJIS.NewEncoder()
				expected, err = encoder.Bytes([]byte(tt.expected))
				if !assertions.NoError(err) {
					return
				}
			} else {
				expected = nil
			}

			assertions.EqualValues(expected, iw.Bytes())
		})
	}
}
func TestNewUTF8WithBOMWriter(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    [][]string
	}{
		{
			"正常系",
			"ヘッダー１,ヘッダー２,ヘッダー３\r\n" +
				"値1,値2,値3\r\n",
			[][]string{
				{
					"ヘッダー１",
					"ヘッダー２",
					"ヘッダー３",
				},
				{
					"値1",
					"値2",
					"値3",
				},
			},
		},
		{
			"empty",
			"",
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			iw := bytes.Buffer{}
			writer := NewUTF8WithBOMWriter(&iw)
			err := writer.WriteAll(tt.given)
			if !assertions.NoError(err) {
				return
			}
			var expected []byte
			if len(tt.expected) > 0 {
				expected = append(UTF8BOM[:], tt.expected...)

			} else {
				expected = nil
			}
			assertions.EqualValues(expected, iw.Bytes())
		})
	}
}
