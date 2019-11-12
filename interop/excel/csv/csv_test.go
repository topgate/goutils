package csv

import (
	"encoding/csv"
	"io"
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
	"github.com/topgate/goutils/reflect"
	"golang.org/x/text/encoding/japanese"
)

func TestReaders(t *testing.T) {
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
			"最終行改行ありCRなし",
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
			"ヘッダー１,ヘッダー２,ヘッダー３\n" +
				"値1,値2,値3\n",
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
			"最終行改行なしCRなし",
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
			"ヘッダー１,ヘッダー２,ヘッダー３\n" +
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
			testReader(t, tt.expected, tt.given, NewSJISReader, toSJIS)
			testReader(t, tt.expected, tt.given, NewReader, toSJIS)
			testReader(t, tt.expected, tt.given, NewReader, toUTF8WithBOM)
			testReader(t, tt.expected, tt.given, NewUTF8WithBOMReader, toUTF8WithBOM)
		})
	}
}

func TestWriters(t *testing.T) {
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
			testWriter(t, tt.expected, tt.given, NewUTF8WithBOMWriter, toUTF8WithBOM)
			testWriter(t, tt.expected, tt.given, NewSJISWriter, toSJIS)
		})
	}
}
func testReader(t *testing.T, expected [][]string, given string, generator func(io.Reader) *csv.Reader, strConverter func(string) []byte) {
	t.Run(reflect.GetFunctionName(generator)+":"+reflect.GetFunctionName(strConverter), func(t *testing.T) {
		assertions := assert.New(t)
		reader := generator(bytes.NewReader(strConverter(given)))

		actual, err := reader.ReadAll()
		if !assertions.NoError(err) {
			return
		}
		assertions.EqualValues(expected, actual)

		individualReader := generator(bytes.NewReader(strConverter(given)))

		index := 0
		for line, err := individualReader.Read(); err != io.EOF; line, err = individualReader.Read() {
			if !assertions.NoError(err) {
				break
			}
			assertions.EqualValues(expected[index], line)
			index++
		}
	})

}

func testWriter(t *testing.T, expected string, given [][]string, generator func(io.Writer) *csv.Writer, strConverter func(string) []byte) {

	t.Run(reflect.GetFunctionName(generator)+":"+reflect.GetFunctionName(strConverter), func(t *testing.T) {
		assertions := assert.New(t)
		iw := &bytes.Buffer{}
		writer := generator(iw)
		err := writer.WriteAll(given)
		if !assertions.NoError(err) {
			return
		}
		if len(expected) > 0 {
			expectedBytes := strConverter(expected)
			assertions.EqualValues(expectedBytes, iw.Bytes())
		} else {
			assertions.Empty(iw.Bytes())
		}

	})
}

func toSJIS(str string) []byte {
	encoder := japanese.ShiftJIS.NewEncoder()
	result, err := encoder.Bytes([]byte(str))
	if err != nil {
		panic(err)
	}
	return result
}

func toUTF8WithBOM(str string) []byte {
	return append(UTF8BOM[:], []byte(str)...)
}
