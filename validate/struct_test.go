package validate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructErrorMessage(t *testing.T) {
	type expected struct {
		err      error
		messages []string
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{
			"len_without_fieldname",
			expected{
				errors.New("Nameの長さは5文字でなければなりません"),
				[]string{"Nameの長さは5文字でなければなりません"},
			},
			struct {
				Name string `validate:"len=5"`
			}{},
		},
		{
			"len",
			expected{
				errors.New("名前の長さは5文字でなければなりません"),
				[]string{"名前の長さは5文字でなければなりません"},
			},
			struct {
				Name string `validate:"len=5" fieldname:"名前"`
			}{},
		},
		{
			"len_ok",
			expected{},
			struct {
				Name string `validate:"len=5" fieldname:"名前"`
			}{Name: "あいうえお"},
		},
		{
			"complex_message",
			expected{
				errors.New("IDは1かより大きくなければなりません\n名前の長さは5文字でなければなりません"),
				[]string{"IDは1かより大きくなければなりません", "名前の長さは5文字でなければなりません"},
			},
			struct {
				ID   int    `validate:"min=1"`
				Name string `validate:"len=5" fieldname:"名前"`
			}{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := Struct(tt.given)
			assertions := assert.New(t)
			ve, ok := actual.(*ValidationError)
			if tt.expected.err == nil {
				assertions.NoError(actual)
				assertions.Nil(ve)
				assertions.False(ok)
			} else {
				assertions.NotNil(ve)
				assertions.True(ok)
				var actualMessages []string
				for _, err := range ve.FiledErrors {
					actualMessages = append(actualMessages, err.Error())
				}
				assertions.EqualValues(tt.expected.messages, actualMessages)
				assertions.EqualError(actual, tt.expected.err.Error())
			}
		})
	}
}
