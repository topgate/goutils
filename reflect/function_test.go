package reflect

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunctionName(t *testing.T) {
	type expected struct {
		name string
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{"関数名", expected{
			name: "reflect.GetFunctionName",
		}, GetFunctionName},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := GetFunctionName(tt.given)
			assertions := assert.New(t)
			assertions.EqualValues(tt.expected.name, actual)
		})
	}

}

func TestGetFunctionName_Panic(t *testing.T) {
	type expected struct {
		err error
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{"関数名", expected{
			err: errors.New("The ptr kind is not function type"),
		}, &struct{}{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover().(error)
				assertions := assert.New(t)
				assertions.EqualError(err, tt.expected.err.Error())
			}()
			actual := GetFunctionName(tt.given)
			assert.EqualValues(t, "", actual)
		})
	}

}

func TestGetFunctionFullName_Panic(t *testing.T) {
	type expected struct {
		err error
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{"関数名", expected{
			err: errors.New("The ptr kind is not function type"),
		}, &struct{}{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover().(error)
				assertions := assert.New(t)
				assertions.EqualError(err, tt.expected.err.Error())
			}()
			actual := GetFunctionFullName(tt.given)
			assert.EqualValues(t, "", actual)
		})
	}

}

func TestGetFunctionFullName(t *testing.T) {
	type expected struct {
		name string
	}
	var tests = []struct {
		name     string
		expected expected
		given    interface{}
	}{
		{"関数名", expected{
			name: "github.com/topgate/goutils/reflect.GetFunctionName",
		}, GetFunctionName},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := GetFunctionFullName(tt.given)
			assertions := assert.New(t)
			assertions.EqualValues(tt.expected.name, actual)
		})
	}

}
