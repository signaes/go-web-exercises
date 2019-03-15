package main

import (
	"fmt"
	"io"
	"net/http"
)

type server struct {
	id string
}

func (m server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/user":
		io.WriteString(res, "Hello")
	case "/whoami":
		io.WriteString(res, fmt.Sprintf("server id: %s", m.id))
	}
}

func main() {
	s := server{"main server 1"}
	http.ListenAndServe(":8080", s)
}
