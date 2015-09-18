package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("[1.0] = %v\n", math.Floor(1.0))
	fmt.Printf("[0.7] = %v\n", math.Floor(0.7))
	fmt.Printf("[-0.5] = %v\n", math.Floor(-0.5))
	fmt.Printf("[-2.0] = %v\n", math.Floor(-2.0))
}
