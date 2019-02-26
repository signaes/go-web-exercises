package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Route struct {
	method, uri string
}

type Routes struct {
	list map[string]Route
}

type home Route

var homePosts = 0

func homeForm() string {
	return `
		<form action="/" method="post">
			<input type="submit" value="POST">
		</form>
	`
}

func showPosts(p int) string {
	return `
		<p>
			<strong>
				POSTS = ` + strconv.Itoa(p) + `
			</strong>
		</p>
	`
}

func (h home) respond() string {
	switch h.method {
	case "GET":
		return template("Olá da home: ", homeForm(), showPosts(homePosts))
	case "POST":
		homePosts++
		return template("Post na home: ", homeForm(), showPosts(homePosts))
	}

	return ""
}

var getHome = home{"GET", "/"}
var postHome = home{"POST", "/"}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	routes := Routes{
		list: map[string]Route{
			"GET /":  Route(getHome),
			"POST /": Route(postHome),
		},
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn, routes)
	}
}

func handle(conn net.Conn, routes Routes) {
	defer conn.Close()

	m, uri := request(conn)

	handler, ok := routes.list[fmt.Sprintf("%s %s", m, uri)]
	status := 404
	response := "Page not found"

	if !ok {
		responseHandler(conn, status, response)

		return
	}

	response = home(handler).respond()
	status = 200

	responseHandler(conn, status, response)
}

func request(conn net.Conn) (string, string) {
	var m, uri string
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			m = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
		}

		if ln == "" {
			break
		}

		i++
	}

	return m, uri
}

func template(s ...string) string {
	return `
		<!DOCTYPE html>
		<html lang="pt-BR">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				` + strings.Join(s, " ") + `
			</body>
		</html>
	`
}

func responseHandler(conn net.Conn, status int, content string) {
	fmt.Fprint(conn, fmt.Sprintf("HTTP/1.1 %d OK\r\n", status))
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(content))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, content)
}
