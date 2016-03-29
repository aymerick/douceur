package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aymerick/douceur/inliner"
	"github.com/aymerick/douceur/parser"
)

const (
	// Version is package version
	Version = "0.2.0"
)

// Usage banner for when the command line is invoked improperly
var Usage = func() {
	fmt.Fprintf(os.Stderr, "USAGE: douceur [parse|inline] filename\n")
	flag.PrintDefaults()
}
var flagVersion bool

func init() {
	flag.BoolVar(&flagVersion, "version", false, "Display version")
	flag.Usage = Usage
}

func main() {
	flag.Parse()

	if flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Missing file path")
		flag.Usage()
		os.Exit(1)
	}

	switch args[0] {
	case "parse":
		parseCSS(args[1])
	case "inline":
		inlineCSS(args[1])
	default:
		fmt.Fprintf(os.Stderr, "Unexpected command: %s", args[0])
		flag.Usage()
		os.Exit(1)
	}
}

// parse and display CSS file
func parseCSS(filePath string) {
	input := readFile(filePath)

	stylesheet, err := parser.Parse(string(input))
	if err != nil {
		fmt.Println("Parsing error: ", err)
		os.Exit(1)
	}

	fmt.Println(stylesheet.String())
}

// inlines CSS into HTML and display result
func inlineCSS(filePath string) {
	input := readFile(filePath)

	output, err := inliner.Inline(string(input))
	if err != nil {
		fmt.Println("Inlining error: ", err)
		os.Exit(1)
	}

	fmt.Println(output)
}

func readFile(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to open file: ", filePath, err)
		os.Exit(1)
	}

	return file
}
