package main

import (
	"fmt"
	"strings"

	"github.com/spiegel-im-spiegel/gutil"
)

func main() {
	msgs := []string {
		"Take the Go-lang!",
		"Go 言語で行こう！",
	}
	fmt.Println(strings.Join(msgs, gutil.LineEnding()))
}
