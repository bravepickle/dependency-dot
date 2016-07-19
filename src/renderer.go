// Render entities to output in dot graphical language

package main

import "bytes"
import "fmt"
import "os"
import "path/filepath"

// generate string in dot format from entities
func RenderViewToDotFormat(entities *[]Entity) string {
	output := bytes.NewBufferString("digraph G { \n")

	for _, entity := range *entities {
		output.WriteString(`  Hello -> `)
		output.WriteString(`"`)
		output.WriteString(entity.Name)
		output.WriteString(`"`)
		output.WriteString(";\n")
	}

	output.WriteString(`}`)

	return output.String()
}

func WriteToFile(outputFile string, rendered *string) {
	file, err := os.Create(outputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, `Failed to create file %s: %s`, outputFile, err)
	}

	_, err = file.WriteString(*rendered)
	if err != nil {
		fmt.Fprintf(os.Stderr, `Failed to write to file %s: %s`, outputFile, err)
	}

	fullpath, err := filepath.Abs(outputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Fprintf(os.Stdin, "Saved to file: %s\n", fullpath)
}
