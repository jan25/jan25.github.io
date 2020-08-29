package main

import (
	"os"
	"text/template"
)

// type Project struct {
// 	Title string
// 	Desc  string
// 	Link  string
// }

// Data is wrapper type to specifiy data variables
// which are injected in templates
type Data struct {
	Title string
}

func main() {
	tmpl := template.Must(template.ParseFiles("about_tmpl.html"))
	data := Data{
		Title: "Abhilash Gnan",
	}
	tmpl.Execute(os.Stdout, data)
}
