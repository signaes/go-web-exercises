package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
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
	name := req.FormValue("name")
	lastName := req.FormValue("lastname")
	subscribe := req.FormValue("subscribe") == "on"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
			<label for="name">
				Name
				<input type="text" name="name">
			</label>
			<label for="lastname">
				Last Name
				<input type="text" name="lastname">
			</label>
			<label for="subscribe">
				Subscribe
				<input type="checkbox" name="subscribe">
			</label>
			<input type="submit">
		</form>
		<br>
	`+name+`<br>`+lastName+`<br>`+strconv.FormatBool(subscribe))
}
