package main

import "fmt"
import "os"

func main() {
	msgs := []string {
		"Take the Go-lang!",
		"Go 言語で行こう！",
	}
	for _, msg := range msgs {
		fmt.Fprintf(os.Stdout, "%s\r\n", msg)
	}
	//fmt.Println("Take the Go-lang!\r\nGo 言語で行こう！")
}
