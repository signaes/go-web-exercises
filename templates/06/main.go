package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	authors := []string{"Donovan", "Kerninghan", "Kennedy", "Ketelsen", "Martin"}

	err := tpl.Execute(os.Stdout, authors)
	if err != nil {
		log.Fatalln(err)
	}
}
