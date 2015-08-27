package text

import (
	"bufio"
	"fmt"
	"io"

	"golang.org/x/text/encoding/japanese" //code.google.com/p/go.text/encoding/japanese
	"golang.org/x/text/transform"         //code.google.com/p/go.text/transform
)

//Conversion
func Conversion(inStream io.Reader, outStream io.Writer) error {
	//read from stream (Shift-JIS to UTF-8)
	scanner := bufio.NewScanner(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder()))
	list := make([]string, 0)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	//write to stream (UTF-8 to EUC-JP)
	writer := bufio.NewWriter(transform.NewWriter(outStream, japanese.EUCJP.NewEncoder()))
	for _, line := range list {
		if _, err := fmt.Fprintln(writer, line); err != nil {
			return err
		}
	}
	return writer.Flush()
}
