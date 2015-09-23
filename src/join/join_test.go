package main

import (
	"os"
	"testing"
)

func readFile() []string {
	file, err := os.Open("CollisionsForHashFunctions.txt") //maybe file path
	if err != nil {
		panic(err)
	}
	defer file.Close()
	list, err := ContentText(file)
	if err != nil {
		panic(err)
	}
	return list
}

func BenchmarkWriteBuffer1(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1(list)
		_ = string(content)
	}
}

func BenchmarkWriteBuffer1Cap128(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1Cap128(list)
		_ = string(content)
	}
}

func BenchmarkWriteBuffer1Cap1K(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer1Cap1K(list)
		_ = string(content)
	}
}

func BenchmarkWriteBuffer2(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2(list)
		_ = string(content)
	}
}

func BenchmarkWriteBuffer2Cap128(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2Cap128(list)
		_ = string(content)
	}
}

func BenchmarkWriteBuffer2Cap1K(b *testing.B) {
	list := readFile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		content := WriteBuffer2Cap1K(list)
		_ = string(content)
	}
}
