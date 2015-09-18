package main

import "fmt"

func main() {
	nihongo := "日本語"
	size := len(nihongo)

	fmt.Printf("nihongo = %d bytes :", size)
	for i := 0; i < size; i++ {
		fmt.Printf(" %x", nihongo[i])
	}
	fmt.Print("\n")
}
