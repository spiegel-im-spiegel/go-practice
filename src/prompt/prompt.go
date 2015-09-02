package main

import (
	"os"
	"strings"

	"github.com/spiegel-im-spiegel/gutil"
)

func main() {
	ui := gutil.CliUi{Reader: os.Stdin, Writer: os.Stdout, ErrorWriter: os.Stderr}
	ui.ModeInteract()
	said := "Hello!"
	err := error(nil)
	for strings.ToUpper(said) != "BYE" {
		said, err = ui.Prompt(said+"> ")
		if err != nil {
			ui.OutputErrln(err)
			break
		}
	}
}
