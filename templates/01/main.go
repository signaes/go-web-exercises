package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	html := template("OK OK")

	fmt.Println(html)

	file, err := os.Create("index.html")
	defer file.Close()
	if err != nil {
		log.Fatal("Error creating file", err)
	}

	io.Copy(file, strings.NewReader(html))
}

func template(content string) string {
	return tabsToSpaces(`
		<!DOCTYPE html>
			<html>
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
				</head>
				` + body(content) + `
			</html>
		</html>
	`)
}

func body(content string) string {
	response := `<body>`
	response += "\n"
	response += `					` + content
	response += "\n"
	response += `				</body>`

	return response
}

func tabsToSpaces(s string) string {
	return strings.Replace(s, "\t", "  ", -1)
}
