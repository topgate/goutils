package gocsv_test

import (
	"os"

	"github.com/topgate/goutils/interop/excel/gocsv"
)

func ExampleMarshal() {
	type User struct {
		FirstName string `csv:"first_name"`
		LastName  string `csv:"last_name"`
	}
	users := []User{
		{
			FirstName: "名前",
			LastName:  "名字",
		},
	}
	file, err := os.Create("user.csv")
	if err != nil {
		panic(err)
	}
	if err := gocsv.Marshal(users, file); err != nil {
		panic(err)
	}
}

func ExampleUnmarshal() {
	type User struct {
		FirstName string `csv:"first_name"`
		LastName  string `csv:"last_name"`
	}
	users := []User{}
	file, err := os.Open("user.csv")
	if err != nil {
		panic(err)
	}

	if err := gocsv.Unmarshal(file, users); err != nil {
		panic(err)
	}
}
