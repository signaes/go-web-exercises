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
	authors := map[string]string{
		"Alan A. A.": "Donovan",
		"Brian W.":   "Kerninghan",
		"William":    "Kennedy",
		"Brian":      "Ketelsen",
		"Erik St.":   "Martin",
	}

	err := tpl.Execute(os.Stdout, authors)
	if err != nil {
		log.Fatalln(err)
	}
}
