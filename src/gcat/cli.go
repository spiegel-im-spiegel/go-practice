package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"gcat/internal/util"
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
		out  string
		text bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&out, "out", "", "output file")
	flags.BoolVar(&text, "text", false, "text mode")

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
	infiles := flags.Args()

	//OutputFile
	if len(out) > 0 {
		file, err := os.Create(out) //maybe file path
		if err != nil {
			fmt.Fprintln(cli.errStream, err)
			return ExitCodeError
		}
		defer file.Close()
		cli.outStream = file
	}

	//Create Context
	con := util.Context {OutStream: cli.outStream, ErrStream: cli.errStream}

	//Input File
	if len(infiles) == 0 {
		con.InStream = cli.inStream
		if err := con.Catenate(text); err != nil {
			fmt.Fprintln(cli.errStream, err)
			return ExitCodeError
		}
	} else {
		for _, infile := range infiles {
			file, err := os.Open(infile) //maybe file path
			if err != nil {
				fmt.Fprintln(cli.errStream, err)
				return ExitCodeError
			}
			defer file.Close()
			con.InStream = file
			if err := con.Catenate(text); err != nil {
				fmt.Fprintln(cli.errStream, err)
				return ExitCodeError
			}
			file.Close()
		}
	}

	return ExitCodeOK
}
