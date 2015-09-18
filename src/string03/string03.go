package main

import "fmt"

func main() {
	nihongo := "日本語"
	nihongoRune := []rune(nihongo)
	size := len(nihongoRune)

	fmt.Printf("nihongo = %d characters : ", size)
	for i := 0; i < size; i++ {
		fmt.Printf("%#U ", nihongoRune[i])
	}
	fmt.Printf("\n")
}
