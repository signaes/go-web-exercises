package main

import (
	"fmt"
	"io"
	"net/http"
)

type server struct {
	id string
}

func (s server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi")
}

type api server

func (a api) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, fmt.Sprintf("{ data: { id: '%s' } }", a.id))
}

func main() {
	s := server{"main server"}
	a := api{"api server"}

	http.Handle("/", s)
	http.Handle("/api", a)

	http.ListenAndServe(":8080", nil)
}
