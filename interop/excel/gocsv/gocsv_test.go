package gocsv

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/topgate/goutils/interop/excel/csv"
)

type Sample struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
}

func TestUnmarshal(t *testing.T) {
	type expected struct {
		err     error
		samples []Sample
	}
	var tests = []struct {
		name     string
		expected expected
		given    io.Reader
	}{
		{
			"Sample1",
			expected{
				samples: []Sample{{Name: "sample_name", ID: 22}},
			},
			strings.NewReader("name,id\nsample_name,22"),
		},
		{
			"無効なCSV",
			expected{
				err:     errors.New("empty csv file given"),
				samples: []Sample{},
			},
			bytes.NewReader(nil),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assertions := assert.New(t)
			actual := []Sample{}
			err := Unmarshal(tt.given, &actual)
			if tt.expected.err != nil {
				assertions.EqualError(err, tt.expected.err.Error())
			} else {
				assertions.NoError(err)
			}
			assertions.EqualValues(tt.expected.samples, actual)
		})
	}

}

func TestMarshal(t *testing.T) {
	type expected struct {
		err  error
		data []byte
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{
			"Sample1",
			expected{
				data: append(csv.UTF8BOM[:], []byte("id,name\r\n22,sample_name\r\n")...),
			},
			[]Sample{
				{
					Name: "sample_name",
					ID:   22,
				},
			},
		},
		{
			"非配列",
			expected{
				err: errors.New("cannot use gocsv.Sample, only slice or array supported"),
			},
			Sample{
				Name: "sample_name",
				ID:   22,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			err := Marshal(tt.given, buffer)
			assertions := assert.New(t)
			if tt.expected.err != nil {
				assertions.EqualError(err, tt.expected.err.Error())
			} else {
				assertions.NoError(err)
			}
			assertions.EqualValues(tt.expected.data, buffer.Bytes())
		})
	}
}
