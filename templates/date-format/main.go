package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func dashToSpaces(s string) string {
	return strings.Replace(s, "-", " ", -1)
}

var fm = template.FuncMap{
	"fmtDate": dayMonthYear,
	"toSpc":   dashToSpaces,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
