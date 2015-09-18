package main

import "fmt"

func main() {
	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	y := 1
	m := 1
	if month < 3 {
		y = year - 1
		m = month + 9
	} else {
		y = year
		m = month - 3
	}
	d := day - 1
	fmt.Printf("y = %d\n", y)
	fmt.Printf("m = %d\n", m)
	fmt.Printf("d = %d\n\n", d)

	mjd := (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
