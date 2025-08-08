package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

	// loop for retry
	for {
		var err error

		// not yet connected or retry by error
		if conn == nil {
			// initialize conn
			conn, err := net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// create a POST request
		request, err := http.NewRequest(
			"POST",
			"http://localhost:8888",
			strings.NewReader(sendMessages[current]),
		)
		if err != nil {
			panic(err)
		}

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		// read from server
		// retry when timeout occurs
		response, err := http.ReadResponse(
			bufio.NewReader(conn),
			request,
		)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		// display result
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		// finish if all has been sent
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
