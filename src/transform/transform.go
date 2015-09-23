package main

import (
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	msgs := []string {
		"Take the Go-lang!",
		"Go 言語で行こう！",
	}
	fmt.Println(strings.Join(msgs, gutil.LineEnding()))
}

func NewDecoder(reader io.Reader) io.Reader {
	return transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder())
}
