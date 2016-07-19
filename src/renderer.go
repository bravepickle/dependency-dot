// Render entities to output in dot graphical language

package main

import "bytes"
import "fmt"
import "os"
import "path/filepath"

// generate string in dot format from entities
func RenderViewToDotFormat(entities *[]Entity) string {
	output := bytes.NewBufferString("digraph G { \n")

	//	rootNodes := getRootEntities(entities)

	for _, entity := range *entities {
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

		fmt.Println(`Children±`, entity.Children)

		for _, child := range entity.Children {
			fmt.Println(`Child`, child.Id)
			output.WriteString(`  `) // padding
			output.WriteString(`"`)
			output.WriteString(entity.Name)
			output.WriteString(`" -> `)
			output.WriteString(`"`)
			output.WriteString(child.Name)
			output.WriteString(`"`)
			output.WriteString(";\n")
		}
	}

	output.WriteString(`}`)

	//	for _, entity := range *rootNodes {
	//		output.WriteString(`  Hello -> `)
	//		output.WriteString(`"`)
	//		output.WriteString(entity.Name)
	//		output.WriteString(`"`)
	//		output.WriteString(";\n")
	//	}

	return output.String()
}

// get all root nodes from entities list
//func getRootEntities(entities *[]Entity) *[]Entity {
//	var rootNodes *[]Entity
//	for _, entity := range *entities {
//		if entity.ParentNode == nil {
//			rootNodes = append(rootNodes, &entity)
//		}
//	}

//	return rootNodes
//}

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
