package validate_test

import (
	"fmt"

	"github.com/topgate/goutils/validate"
)

func Example() {
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

func Example_detailError() {

	type User struct {
		ID   int    `validate:"min=1"`
		Name string `validate:"len=5" fieldname:"名前"`
	}

	user := User{
		ID:   0,
		Name: "花山太郎",
	}

	err := validate.Struct(user)
	errs, _ := validate.GetValidationErrors(err)
	fmt.Println(errs[0].Translate(validate.Translator))
	fmt.Println(errs[1].Translate(validate.Translator))
	// Output:
	// IDは1かより大きくなければなりません
	// 名前の長さは5文字でなければなりません

}
