package util

import (
	"os"
	"testing"
)

var list []string

func readFile() {
	file, err := os.Open("testdata/CollisionsForHashFunctions.txt") //maybe file path
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list, err = ContentText(file)
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteBuffer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1(list)
		_ = content
	}
}

func BenchmarkWriteBuffer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2(list)
		_ = content
	}
}
