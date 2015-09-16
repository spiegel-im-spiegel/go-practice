package main

import "fmt"

func main() {
	//a := make([]int, 0)                                         // 空の配列を用意
	a := make([]int, 0, 32)                                     // 空の配列を用意
	fmt.Printf("Slice(%02d) : %p : %v (%v)\n", 0, a, a, cap(a)) // 配列の表示（初期状態）
	for num := 1; num <= 17; num++ {
		a = append(a, num)                                            //配列要素の追加
		fmt.Printf("Slice(%02d) : %p : %v (%v)\n", num, a, a, cap(a)) //配列の表示
	}
}
