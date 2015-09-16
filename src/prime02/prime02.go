package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	primes := make([]int64, 1)     // 素数のリスト
	primes_f := make([]float64, 1) // 素数のリスト（浮動小数点へのキャスト）
	primes[0] = 2                  // 2 は素数
	primes_f[0] = 2.0              // 2 は素数（浮動小数点）
	var max int64 = 100

	start := time.Now() // Start
	var n int64 = 3
	for n = 3; n < max; n += 2 { // 3 から始まる奇数のみを探索
		flag := true
		f := float64(n)                    // 浮動小数点に cating
		rf := math.Sqrt(f)                 // n に対して √n をとる
		for i := 1; i < len(primes); i++ { // 2 より大きい既知の素数でチェックする
			if primes_f[i] > rf { // n に対して √n 以下の素数まで探索すればよい
				break
			} else if (n % primes[i]) == 0 { // n が既知の素数で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)     // 素数を追加
			primes_f = append(primes_f, f) // 素数を追加（浮動小数点）
		}
	}
	goal := time.Now() // Goal
	fmt.Printf("%v 以下の素数: %v\n", max, primes)
	fmt.Printf("%v 経過", goal.Sub(start)) // 経過時間を表示
}
