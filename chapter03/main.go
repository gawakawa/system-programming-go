package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	teeReader()
}

func read(r io.Reader) {
	buffer := make([]byte, 1024)
	size, err := r.Read(buffer)
	_ = size
	_ = err
}

func read_all(r io.Reader) {
	buffer, err := io.ReadAll(r)
	_ = buffer
	_ = err
}

func read_full(r io.Reader) {
	buffer := make([]byte, 4)
	size, err := io.ReadFull(r, buffer)
	_ = size
	_ = err
}

func copy(w io.Writer, r io.Reader) {
	writeSize, err := io.Copy(w, r)
	_ = writeSize
	_ = err
}

func copyN(w io.Writer, r io.Reader, size int64) {
	writeSize, err := io.CopyN(w, r, size)
	_ = writeSize
	_ = err
}

func cast(reader io.Reader, writer io.Writer) {
	var readCloser io.ReadCloser = io.NopCloser(reader)
	var readWriter io.ReadWriter = bufio.NewReadWriter(bufio.NewReader(reader), bufio.NewWriter(writer))
	_ = readCloser
	_ = readWriter
}

func ReadStdin() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

func ReadFile() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func ReadNet() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: example.com\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	fmt.Println(res.Header)
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func SectionReader() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	SectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, SectionReader)
}

func ConvertEndian() {
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func readChunks(file *os.File) []io.Reader {
	// array holding chunks
	var chunks []io.Reader

	// skip leading 8 bytes
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// move to a head of a next chunk
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	writer := io.MultiWriter(&buffer, crc)
	io.WriteString(writer, "teXt")
	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func embedSecretText() {
	file, err := os.Open("PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	// write signature
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// write a IHDR chunk
	io.Copy(newFile, chunks[0])
	// add a text chunk
	io.Copy(newFile, textChunk("Lambda Note++"))
	// add rest chunks
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func dumpChunks() {
	file, err := os.Open("PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func readText() {
	var source = `1行め
2行め
3行め`

	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}

func readScanner() {
	var source = `1行め
2行め
3行め`

	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

func fscan() {
	var source = "123 1.234 1.0e4 test"

	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}

func readCsv(csvSource string) {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}

func multiReader() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.multiReader\n")
	footer := bytes.NewBufferString("----- FOOTER -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}

func teeReader() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = io.ReadAll(teeReader)
	fmt.Println(buffer.String())
}
