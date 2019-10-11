package csv

import (
	"testing"

	"bytes"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/japanese"
)

func TestNewReader(t *testing.T) {

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
			actual, err := reader.ReadAll()
			if !assertions.NoError(err) {
				return
			}
			assertions.EqualValues(tt.expected, actual)
		})
	}
}

func TestNewWriter(t *testing.T) {
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
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			iw := bytes.Buffer{}
			writer := NewWriter(&iw)
			err := writer.WriteAll(tt.given)
			if !assertions.NoError(err) {
				return
			}
			encoder := japanese.ShiftJIS.NewEncoder()
			expected, err := encoder.Bytes([]byte(tt.expected))
			if !assertions.NoError(err) {
				return
			}

			assertions.EqualValues(expected, iw.Bytes())
		})
	}
}
