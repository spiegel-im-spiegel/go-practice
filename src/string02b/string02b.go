package main

import "fmt"

func main() {
	nihongo := "日本語"

	fmt.Printf("nihongo = %s\n", nihongo)
	fmt.Printf("nippon = %s\n", string([]rune(nihongo)[:2]))
}
