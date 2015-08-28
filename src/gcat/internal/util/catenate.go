package util

import (
	"bytes"
	"bufio"
	//"fmt"
	"io"
	//"os"
)

// Context of catenateRaw
type Context struct {
	// inStream is the stdin.
	InStream io.Reader

	// outStream and errStream are the stdout and stderr
	OutStream, ErrStream io.Writer
}

//Catenation (raw data)
func (con *Context) Catenate(text bool) error {
	if text {
		return con.CatenateText()
	} else {
		return con.CatenateRaw()
	}
}

//Catenation (raw data)
func (con *Context) CatenateRaw() error {
	//read from stream
	content, err := ContentRaw(con.InStream)
	if err != nil {
		return err
	}
	//write to stream
	writer := bufio.NewWriter(con.OutStream)
	if _, err := writer.Write(content); err != nil {
		return err
	}
	return writer.Flush()
}

//Read content (raw data) from buffer
func ContentRaw(inStream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(inStream); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//Catenation (text data)
func (con *Context) CatenateText() error {
	//read from stream
	lines, err := ContentText(con.InStream)
	if err != nil {
		return err
	}

	//write to byte buffer (temporary)
	content := WriteBuffer1(lines)
	//content := WriteBuffer2(lines)

	//write to stream
	writer := bufio.NewWriter(con.OutStream)
	if _, err := writer.Write(content); err != nil {
		return err
	}
	return writer.Flush()
}

//Write content (text data) to buffer
func ContentText(inStream io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(inStream)
	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

//Read content (text data) from buffer
func WriteBuffer1(lines []string) []byte {
	//write to byte buffer
	content := make([]byte, 0)
	//content := make([]byte, 0, 1024) //1K bytes capacity
	recode := "\r\n"
	//fmt.Fprintf(os.Stderr, "%v (%p)\n", cap(content), content)
	for _, line := range lines {
		content = append(content, line...)
		content = append(content, recode...)
		//fmt.Fprintf(os.Stderr, "%v (%p)\n", cap(content), content)
	}
	return content
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2(lines []string) []byte {
	//write to byte buffer
	content := bytes.NewBuffer(make([]byte, 0))
	//content := bytes.NewBuffer(make([]byte, 0, 1024)) //1K bytes capacity
	recode := "\r\n"
	for _, line := range lines {
		content.WriteString(line)
		content.WriteString(recode)
	}
	return content.Bytes()
}
