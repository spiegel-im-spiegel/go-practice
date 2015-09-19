package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ModifiedJulianDayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return mjdnUnix(t)
	} else {
		return mjdnGregorian(t)
	}
}

func mjdnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func mjdnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Fprintln(os.Stderr, "年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		} else {
			args[i] = num
		}
	}
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", ModifiedJulianDayNumber(tm))
}
