package main

import "fmt"

func main() {
	err := r()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Normal End.")
	}
}

func r() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("Recovered from: %v", rec)
		}
	}()

	f()
	err = nil
	return
}

func f() {
	panic("Panic!")
}
