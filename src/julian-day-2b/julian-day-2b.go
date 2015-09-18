package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int64, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.ParseInt(argsStr[i], 10, 64)
		if err != nil {
			if enum, ok := err.(*strconv.NumError); ok {
				switch enum.Err {
				case strconv.ErrRange:
					fmt.Printf("Bad Range Error")
				case strconv.ErrSyntax:
					fmt.Printf("Syntax Error")
				}
			}
			return
		} else {
			args[i] = num
		}
	}
	fmt.Printf("%v年%v月%v日\n\n", args[0], args[1], args[2])

	y := args[0]
	m := args[1]
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := args[2] - 1
	fmt.Printf("y = %d\n", y)
	fmt.Printf("m = %d\n", m)
	fmt.Printf("d = %d\n\n", d)

	mjd := (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
