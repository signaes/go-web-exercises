package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(template("OK OK"))
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
