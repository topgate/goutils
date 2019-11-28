package validate_test

import (
	"fmt"

	"github.com/topgate/goutils/validate"
)

func ExampleStruct() {
	type User struct {
		ID   int    `validate:"min=1"`
		Name string `validate:"len=5" fieldname:"名前"`
	}

	user := User{
		ID:   0,
		Name: "花山太郎",
	}

	err := validate.Struct(user)
	fmt.Print(err.Error())
	// Output:
	// IDは1かより大きくなければなりません
	// 名前の長さは5文字でなければなりません
}

func ExampleStruct_fieldErrors() {

	type User struct {
		ID   int    `validate:"min=1"`
		Name string `validate:"len=5" fieldname:"名前"`
	}

	user := User{
		ID:   0,
		Name: "花山太郎",
	}

	err := validate.Struct(user)
	ve, ok := err.(*validate.ValidationError)
	if ok {
		fmt.Println(ve.FiledErrors[0].Error())
		fmt.Println(ve.FiledErrors[1].Error())
	}
	// Output:
	// IDは1かより大きくなければなりません
	// 名前の長さは5文字でなければなりません

}
