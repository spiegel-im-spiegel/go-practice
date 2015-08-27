package binary

import (
	"bytes"
	"bufio"
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//Conversion
func Conversion(inStream io.Reader, outStream io.Writer) error {
	//read from stream (Shift-JIS to UTF-8)
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder())); err != nil {
		return err
	}

	//write to stream (UTF-8 to EUC-JP)
	writer := bufio.NewWriter(transform.NewWriter(outStream, japanese.EUCJP.NewEncoder()))
	if _, err := writer.Write(buf.Bytes()); err != nil {
		return err
	}
	return writer.Flush()
}
