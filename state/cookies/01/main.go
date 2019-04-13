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

	if err != nil {
		counter = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	} else {
		strValue := counter.Value

		if strValue != "" {
			value, err = strconv.Atoi(strValue)

			if err != nil {
				value = 0
			}

			value++
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(value),
	})

	io.WriteString(w, fmt.Sprintf("Hello: %d", value))
}
