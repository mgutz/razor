package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mgutz/razor/razor"
)

var configFile string
var debug bool

func init() {
	flag.StringVar(&configFile, "config", "razor.yml", "YAML config filename")
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
	razor.Init(configFile, debug)
	options := razor.Option{}

	if debug {
		options["Debug"] = true
	}
	if flag.NArg() >= 2 {
		arg1, arg2 := flag.Arg(0), flag.Arg(1)
		stat, err := os.Stat(arg1)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if stat.IsDir() {
			fmt.Printf("Gorazor processing dir: %s -> %s\n", arg1, arg2)
			err := razor.GenFolder(arg1, arg2, options)
			if err != nil {
				fmt.Println(err)
			}
		} else if stat.Mode().IsRegular() {
			fmt.Printf("Gorazor processing file: %s -> %s\n", arg1, arg2)
			razor.GenFile(arg1, arg2, options)
		} else {
			flag.Usage()
		}
	}
}
