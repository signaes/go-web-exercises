package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

type person struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var people = map[string]person{}
var sessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/options", options)
	http.Handle("favicon.icon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sID, e := uuid.NewV4()
		if e != nil {
			log.Fatal(e)
		}

		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}

	var p person
	if pID, ok := sessions[c.Value]; ok {
		p = people[pID]
	}

	if req.Method == http.MethodPost {
		pID := req.FormValue("username")
		firstName := req.FormValue("firstname")
		lastName := req.FormValue("lastname")
		p = person{pID, firstName, lastName}
		sessions[c.Value] = pID
		people[pID] = p
	}

	tpl.ExecuteTemplate(res, "index.gohtml", p)
}

func options(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	pID, ok := sessions[c.Value]
	if !ok {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	p := people[pID]
	tpl.ExecuteTemplate(res, "options.gohtml", p)
}
