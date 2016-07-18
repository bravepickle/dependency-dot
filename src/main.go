// Entry point for dependency-dot application
// Author: Victor K
// License: MIT

package main

import (
	"flag"
	"fmt"
	"os"
)

var nameTitle = flag.String("n", "name", "Entity name column")
var refTitle = flag.String("r", "reference", "Column that contains references to other entities. References must be comma-separated, e.g. '1,14,22'")
var idTitle = flag.String("i", "id", "Column that contains entity IDs")
var outputDirTitle = flag.String("d", ".", "Output directory")
var verbose = flag.Bool("v", false, "Increase verbosity in output")

// show help info on usage and finish application
func showHelp() {
	fmt.Fprintln(os.Stderr, `Generate dot notation from CSV file. First row should contain column titles`)
	fmt.Fprint(os.Stderr, os.Args[0], `-i [ID]`, `-r [REFS]`, `input.csv`, "\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\nExamples:\n  %s -i ID -r References -n Names data.csv\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -d ../output data.csv\n", os.Args[0])
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		showHelp()
	}

	input := flag.Arg(0)
	Debug = *verbose
	if Debug {
		fmt.Printf("Input:\n  Input Filename = %s\n  ID Title = %s\n  Name Title = %s\n  Reference Title = %s\n  Output Directory = %s\n", flag.Arg(0), *idTitle, *nameTitle, *refTitle, *outputDirTitle)
	}

	columns := map[string]string{`id`: *idTitle, `name`: *nameTitle, `ref`: *refTitle}

	ParseCsv(input, columns)
}
