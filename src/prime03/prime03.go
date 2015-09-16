package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	//コマンドライン引数の解析
	algno := flag.Int("alg", 0, "0: Basic algorithm , 1: Sieve of Eratosthenes")
	flag.Parse()
	args := flag.Args()
	if *algno < 0 || *algno > 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	max, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if max <= 0 {
		max = 1
	}

	//素数探索
	prime := int64(0)
	start := time.Now() // Start
	switch *algno {
	case 1:
		prime = LastPrimeE(max)
	default:
		prime = LastPrimeB(max)
	}
	goal := time.Now()                       // Goal
	fmt.Printf("%v 個目の素数: %v\n", max, prime) // max 個目の素数
	fmt.Printf("%v 経過\n", goal.Sub(start))   // 経過時間を表示
}

func LastPrimeB(max int64) int64 {
	count := int64(0)

	for n := int64(2); ; n++ {
		flag := true
		for m := int64(2); m < n; m++ {
			if (n % m) == 0 { // n が m で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
		}
	}
}

func LastPrimeE(max int64) int64 {
	if max <= 1 {
		return 2
	}
	primes := make([]int64, 1, max)     // 素数のリスト
	primes_f := make([]float64, 1, max) // 素数のリスト（浮動小数点へのキャスト）
	primes[0] = 2                       // 2 は素数
	primes_f[0] = 2.0                   // 2 は素数（浮動小数点）

	count := int64(1)
	for n := int64(3); ; n += 2 { // 3 から始まる奇数のみを探索
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
			count++
			if count >= max {
				return n
			}
			primes = append(primes, n)     // 素数を追加
			primes_f = append(primes_f, f) // 素数を追加（浮動小数点）
		}
	}
}
