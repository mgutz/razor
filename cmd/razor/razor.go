package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mgutz/razor"
)

var configFile string
var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "print debug info")
}

func Usage() {
	fmt.Fprintf(os.Stderr, "usage: gorazor <input dir or file> <output dir or file> [-debug]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	options := razor.Option{}

	if debug {
		options["Debug"] = true
	}
	if flag.NArg() >= 2 {
		var err error
		arg1, arg2 := flag.Arg(0), flag.Arg(1)
		orig1, orig2 := arg1, arg2

		if arg1 == "." {
			arg1, err = filepath.Abs(arg1)
			if err != nil {
				fmt.Errorf("Could not convert to absolute path: \"%s\"\n%v\n", arg1, err)
				os.Exit(2)
			}
		}

		if arg2 == "." {
			arg2, err = filepath.Abs(arg2)
			if err != nil {
				fmt.Errorf("Could not convert to absolute path: \"%s\"\n%v\n", arg2, err)
				os.Exit(2)
			}
		}

		stat, err := os.Stat(arg1)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if stat.IsDir() {
			fmt.Printf("Gorazor processing dir: %s -> %s\n", orig1, orig2)
			err := razor.GenFolder(arg1, arg2, options)
			if err != nil {
				fmt.Println(err)
			}
		} else if stat.Mode().IsRegular() {
			fmt.Printf("Gorazor processing file: %s -> %s\n", orig1, orig2)
			razor.GenFile(arg1, arg2, options)
		} else {
			flag.Usage()
		}
	}
}
