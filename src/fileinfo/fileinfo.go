package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var err error

	//open file
	var file *os.File
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 1 {
		file = os.Stdin
	} else {
		file, err = os.Open(argsStr[0]) //maybe file path
		if err != nil {
			fmt.Fprintln(os.Stderr, err) //error message
			return
		}
		defer file.Close() //close file at return
	}

	//file status
	var stat os.FileInfo
	stat, err = file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err) //error message
		return
	}
	fmt.Fprintln(os.Stdout, "    file name : ", stat.Name())
	fmt.Fprintln(os.Stdout, "    file size : ", stat.Size())
	fmt.Fprintln(os.Stdout, "   permission : ", stat.Mode())
	fmt.Fprintln(os.Stdout, "    timestamp : ", stat.ModTime())
	fmt.Fprintln(os.Stdout, "Is Directory? : ", stat.IsDir())

	return
}
