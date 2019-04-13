package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon/ico", http.NotFoundHandler())
	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	var value int
	counter, err := req.Cookie("counter")

	if err == http.ErrNoCookie {
		counter = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	value, _ = strconv.Atoi(counter.Value)
	value++
	counter.Value = strconv.Itoa(value)

	http.SetCookie(w, counter)
	io.WriteString(w, fmt.Sprintf("Hello: %d", value))
}
