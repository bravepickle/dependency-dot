// Render entities to output in dot graphical language

package main

import "bytes"
import "fmt"
import "os"
import "path/filepath"

// generate string in dot format from entities
func RenderViewToDotFormat(entities *[]Entity, appendDot string) string {
	output := bytes.NewBufferString("digraph G { \n")

	if appendDot != `` {
		output.WriteString(`  `) // padding
		output.WriteString(appendDot)
		output.WriteString("\n")
	}

	for _, entity := range *entities {
		addNode(output, entity)
		//		fmt.Println(`Children±`, entity.Children)

		for _, child := range entity.Children {
			output.WriteString(`  `) // padding
			output.WriteString(`"`)
			output.WriteString(entity.Name)
			output.WriteString(`" -> `)
			output.WriteString(`"`)
			output.WriteString(child.Name)
			output.WriteString(`"`)

			ref, ok := entity.Ref(child.Id)

			if ok && ref.Style != `` {
				output.WriteString(` [`)
				output.WriteString(ref.Style)
				output.WriteString(`]`)
			}

			output.WriteString(";\n")
		}
	}

	output.WriteString(`}`)

	return output.String()
}

// add node definition
func addNode(output *bytes.Buffer, entity Entity) {
	output.WriteString(`  `) // padding
	output.WriteString(`"`)
	output.WriteString(entity.Name)
	output.WriteString(`"`)

	if entity.Style != `` {
		output.WriteString(` [`)
		output.WriteString(entity.Style)
		output.WriteString(`]`)
	}

	output.WriteString(";\n")
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

	fmt.Fprintf(os.Stdout, "Saved to file: %s\n", fullpath)
}
