package binary

import (
	"io"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//Conversion
func Conversion(inStream io.Reader, outStream io.Writer) error {
	//reader from stream (Shift-JIS to UTF-8)
	reader := transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder())

	//writer to stream (UTF-8 to EUC-JP)
	writer := transform.NewWriter(outStream, japanese.EUCJP.NewEncoder())

	//Copy
	_, err := io.Copy(writer, reader)
	return err
}
