package main

import "fmt"

func main() {
	nihongo := "日本語"

	for pos, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, pos)
	}
}
