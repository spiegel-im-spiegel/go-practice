package main

import (
	"fmt"
	"time"
)

func main() {
	const onday = int64(86400)   //seconds
	const baseDay = int64(40587) //Modified Julian Date at January 1, 1970

	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	tm := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	ut := tm.Unix()
	fmt.Printf("UNIX Time = %v seconds = %v days and %v seconds\n\n", ut, ut/onday, ut%onday)

	mjd := ut/onday + baseDay
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
