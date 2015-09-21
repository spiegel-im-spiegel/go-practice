package main

import (
	"bufio"
	"bytes"
	"io"
)

//Read content (text data) from buffer
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

//Write content (text data) to buffer
func WriteBuffer1(lines []string) []byte {
	//write to byte buffer
	content := make([]byte, 0)
	recode := "\r\n"
	for _, line := range lines {
		content = append(content, line...)
		content = append(content, recode...)
	}
	return content
}

//Write content (text data) to buffer
func WriteBuffer1Cap128(lines []string) []byte {
	//write to byte buffer
	content := make([]byte, 0, 128) //128 bytes capacity
	recode := "\r\n"
	for _, line := range lines {
		content = append(content, line...)
		content = append(content, recode...)
	}
	return content
}

//Write content (text data) to buffer
func WriteBuffer1Cap1K(lines []string) []byte {
	//write to byte buffer
	content := make([]byte, 0, 1024) //1K bytes capacity
	recode := "\r\n"
	for _, line := range lines {
		content = append(content, line...)
		content = append(content, recode...)
	}
	return content
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2(lines []string) []byte {
	//write to byte buffer
	content := bytes.NewBuffer(make([]byte, 0))
	recode := "\r\n"
	for _, line := range lines {
		content.WriteString(line)
		content.WriteString(recode)
	}
	return content.Bytes()
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2Cap128(lines []string) []byte {
	//write to byte buffer
	content := bytes.NewBuffer(make([]byte, 0, 128)) //128 bytes capacity
	recode := "\r\n"
	for _, line := range lines {
		content.WriteString(line)
		content.WriteString(recode)
	}
	return content.Bytes()
}

//Write content (text data) to buffer (buffered I/O)
func WriteBuffer2Cap1K(lines []string) []byte {
	//write to byte buffer
	content := bytes.NewBuffer(make([]byte, 0, 1024)) //1K bytes capacity
	recode := "\r\n"
	for _, line := range lines {
		content.WriteString(line)
		content.WriteString(recode)
	}
	return content.Bytes()
}
