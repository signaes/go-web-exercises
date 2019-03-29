package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon/ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func logMiddleware(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("q")
	fmt.Println("Listening at :8080")
	fmt.Println("Searching for " + q)
}

func index(w http.ResponseWriter, req *http.Request) {
	logMiddleware(w, req)
	v := req.FormValue("q")
	io.WriteString(w, "Search: "+v)
}
