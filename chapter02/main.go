package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}

	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}

	gzip_writer := gzip.NewWriter(file)
	gzip_writer.Header.Name = "test.txt"

	writer := io.MultiWriter(gzip_writer, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "	")
	encoder.Encode(source)

	gzip_writer.Flush()
	file.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func write_file() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}

func write_console() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

func write_buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

func write_builder() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}

func write_net() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	// io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)
	io.Copy(os.Stdout, conn)
}

func write_browser_handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func write_browser() {
	http.HandleFunc("/", write_browser_handler)
	http.ListenAndServe(":8080", nil)
}

func multiwriter() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}

func write_zipped_data_to_file() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}

func writer_bufio() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}

func fprintf() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}

func print_json() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}

func write_request() {
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "ヘッダーも追加できます")
	request.Write(os.Stdout)
}

func fprintf_file() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "string: %s\nint: %d\nfloat: %f", "hello", 42, 3.14)
}

func write_csv_to_stdout() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Go", "Rust", "Haskell"})
	writer.Flush()
	writer.Write([]string{"bad", "not bad", "good"})
	writer.Flush()
}

func write_csv_to_file() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"Go", "Rust", "Haskell"})
	writer.Flush()
	writer.Write([]string{"bad", "not bad", "good"})
	writer.Flush()
}
