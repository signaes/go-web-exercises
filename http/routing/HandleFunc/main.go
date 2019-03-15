package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello from HandleFunc")
	})
	http.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "Hello from api")
	})

	http.ListenAndServe(":8080", nil)
}
