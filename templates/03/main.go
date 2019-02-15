package main

import (
	"log"
	"os"
	"text/template"
)

const (
	a = iota
	b
	c
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, a)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "template.gohtml", b)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "other.gohtml", c)
	if err != nil {
		log.Fatalln(err)
	}
}
