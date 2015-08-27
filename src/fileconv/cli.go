package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"fileconv/internal/binary"
	"fileconv/internal/text"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// inStream is the stdin.
	inStream io.Reader

	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		bin bool
		out string //maybe file path
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&bin, "bin", false, "binary mode")
	flags.StringVar(&out, "out", "", "output file")

	flVersion := flags.Bool("version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if *flVersion {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	// Parse commandline sub-arguments
	subArgs := flags.Args()
	inputFile := ""
	if len(subArgs) > 1 {
		fmt.Fprintln(cli.errStream, os.ErrInvalid, subArgs)
		return ExitCodeError
	} else if len(subArgs) == 1 {
		inputFile = subArgs[0] //maybe file path
	}

	//Input File
	var infile io.Reader
	if len(inputFile) == 0 {
		infile = cli.inStream
	} else {
		file, err := os.Open(inputFile) //maybe file path
		if err != nil {
			fmt.Fprintln(cli.errStream, err)
			return ExitCodeError
		}
		defer file.Close()
		infile = file
	}

	//OutputFile
	var outfile io.Writer
	if len(out) == 0 {
		outfile = cli.outStream
	} else {
		file, err := os.Create(out) //maybe file path
		if err != nil {
			fmt.Fprintln(cli.errStream, err)
			return ExitCodeError
		}
		defer file.Close()
		outfile = file
	}

	var err error
	if bin {
		err = binary.Conversion(infile, outfile)
	} else {
		err = text.Conversion(infile, outfile)
	}
	if err != nil {
		fmt.Fprintln(cli.errStream, err)
		return ExitCodeError
	}

	return ExitCodeOK
}
