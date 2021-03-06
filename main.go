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
var refTitle = flag.String("r", "reference", "Column that contains references to other entities. References must be comma-separated, e.g. '1,14,22'. Can contain dot styles wrapped by square brackets")
var idTitle = flag.String("i", "id", "Column that contains entity IDs")
var groupTitle = flag.String("g", "group", "Column that contains groups to group nodes as subgraphs")
var styleTitle = flag.String("s", "style", "Column that contains (optional) styles for nodes in dot language format without square brackets")
var outputFile = flag.String("o", "", "Output file. If not set, then output to STDIN")
var appendDot = flag.String("a", "", "Append custom attributes to dot file, e.g. 'size =\"4,4\";nodesep=1.05;rankdir=LR;'. See dot guide for attributes definition and reference")
var appendDir = flag.String("d", "", "Append custom arrows styling for all, e.g. 'dir=back'")
var verbose = flag.Bool("v", false, "Increase verbosity in output")

// show help info on usage and finish application
func showHelp() {
	fmt.Fprintln(os.Stderr, `Generate dot notation from CSV file. First row should contain column titles`)
	fmt.Fprint(os.Stderr, os.Args[0], ` -i [ID]`, ` -r [REFS]`, ` -n [NAME]`, ` -o [OUTPUT]`, ` -s [STYLE]`, ` -g [GROUP]`, ` -v`, ` input.csv`, "\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\nExamples:\n  %s -i ID -r References -n Names data.csv\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -o ../test_data.dot data.csv\n", os.Args[0])
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
		fmt.Printf("Input:\n  Input Filename = %s\n  ID Title = %s\n  Name Title = %s\n  Reference Title = %s\n  Output File = %s\n  Group Title = %s\n\n", flag.Arg(0), *idTitle, *nameTitle, *refTitle, *outputFile, *groupTitle)
	}

	columns := map[string]string{`id`: *idTitle, `name`: *nameTitle, `ref`: *refTitle, `style`: *styleTitle, `group`: *groupTitle}
	entities := ParseCsv(input, columns)
	rendered := RenderViewToDotFormat(&entities, *appendDot, *appendDir)

	if *outputFile != `` {
		WriteToFile(*outputFile, &rendered)
	} else {
		fmt.Fprintln(os.Stdout, rendered)
	}
}
