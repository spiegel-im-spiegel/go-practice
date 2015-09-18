package main

import (
	"fmt"
	"math"
)

func main() {
	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	mm := float64(month - 3)
	y := float64(year) + math.Floor(mm/12.0)
	m := math.Mod(12.0+mm, 12.0)
	d := float64(day - 1)
	fmt.Printf("y = %f\n", y)
	fmt.Printf("m = %f\n", m)
	fmt.Printf("d = %f\n\n", d)

	mjd := math.Floor(365.25*y) + math.Floor(y/400.0) - math.Floor(y/100.0) + math.Floor(30.60*m+0.5) + d - 678881.0
	fmt.Printf("MJD = %f日\n", mjd)
	fmt.Printf("JD = %f日\n", mjd+2400000.5)
}
