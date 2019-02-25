package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("\tMETHOD =", m)
		} else if i == 1 {
			uri := strings.Fields(ln)[1]
			fmt.Println("\tURI =", uri)
		}

		if ln == "" {
			break
		}

		i++
	}
}

func respond(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="pt-BR">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
				<strong>Olá</strong>
			</body>
		</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
