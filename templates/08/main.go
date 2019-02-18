package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type author struct {
	Name, Surname string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	authors := []author{
		author{"Alan A. A.", "Donovan"},
		author{"Brian W.", "Kerninghan"},
		author{"William", "Kennedy"},
		author{"Brian", "Ketelsen"},
		author{"Erik St.", "Martin"},
	}

	err := tpl.Execute(os.Stdout, authors)
	if err != nil {
		log.Fatalln(err)
	}
}
