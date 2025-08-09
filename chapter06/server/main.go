package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// Can a client accept gzip?
func isGZipAcceptable(request *http.Request) bool {
	return strings.Contains(
		strings.Join(request.Header["Accept-Encoding"], ","),
		"gzip",
	)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// loop for keep-alive
			for {
				// set timeout
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				request, err := http.ReadRequest(
					bufio.NewReader(conn),
				)
				if err != nil {
					neterr, ok := err.(net.Error)
					// timeout
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
						// socket is close
					} else if err == io.EOF {
						break
					}
					panic(err)
				}

				// display request
				dump, err := httputil.DumpRequest(request, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"

				// write response
				response := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body: io.NopCloser(
						strings.NewReader(content),
					),
				}
				response.Write(conn)
			}
		}()
	}
}
