package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Stu Jensen",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
