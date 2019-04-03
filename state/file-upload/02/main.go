package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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
	var s string

	logMiddleware(w, req)
	name := req.FormValue("name")
	lastName := req.FormValue("lastname")
	subscribe := req.FormValue("subscribe") == "on"

	fmt.Println(req.Method)

	if req.Method == http.MethodPost {
		// open file
		f, h, err := req.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader", h, "\nerr", err)

		// read file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// create folder if not exist
		newpath := filepath.Join(".", "public")
		err = os.MkdirAll(newpath, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// store on server
		dst, err := os.Create(filepath.Join("./public/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post" enctype="multipart/form-data">
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
			<label for="file">
				<input type="file" name="file">
			</label>
			<input type="submit">
		</form>
		<br>
	`+name+`<br>`+lastName+`<br>`+strconv.FormatBool(subscribe)+`<br>`+`<h1>File</h1>`+s)
}
