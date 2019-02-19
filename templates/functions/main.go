package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type author struct {
	Name, Surname string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	return strings.TrimSpace(s)[:3]
}

func main() {
	authors := []author{
		author{"Alan A. A.", "Donovan"},
		author{"Brian W.", "Kerninghan"},
		author{"William", "Kennedy"},
		author{"Brian", "Ketelsen"},
		author{"Erik St.", "Martin"},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", authors)
	if err != nil {
		log.Fatalln(err)
	}
}
