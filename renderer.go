// Render entities to output in dot graphical language

package main

import "bytes"
import "fmt"
import "sort"
import "os"
import "path/filepath"
import "strconv"

const defaultPadding = `  `

type groupEntities map[string]entityGroup

type entityGroup struct {
	Id       int
	Name     string
	Entities []Entity
}

// check if current group is default
func (g *entityGroup) IsDefault() bool {
	return g.Name == `` // empty group is default one
}

func (g *entityGroup) Add(e Entity) {
	g.Entities = append(g.Entities, e)
}

func (g *entityGroup) Wrap(output *bytes.Buffer) {
	if g.IsDefault() { // no subgraphs needed
		addGraphNodes(&g.Entities, output, defaultPadding)
	} else {
		padding := defaultPadding + defaultPadding
		output.WriteString(`  subgraph `) // padding
		output.WriteString(`cluster`)
		output.WriteString(strconv.Itoa(g.Id))
		output.WriteString(" {\n")

		output.WriteString(padding)
		output.WriteString(`label = "`)
		output.WriteString(g.Name)
		output.WriteString(`"`)
		output.WriteString(";\n")

		output.WriteString(padding)
		output.WriteString("color=black;\n")

		output.WriteString(padding)
		output.WriteString("style=dashed;\n")

		addGraphNodes(&g.Entities, output, padding)
		output.WriteString("  }\n")
	}
}

func groupEntitiesArr(entities *[]Entity) (grEntities groupEntities) {
	grEntities = make(groupEntities)
	i := 1
	for _, e := range *entities {
		group, ok := grEntities[e.Group]
		//		if _, ok := grEntities[e.Group]; !ok {
		if !ok {
			group = entityGroup{Name: e.Group, Id: i, Entities: []Entity{}}
			i++
		}

		group.Add(e)
		grEntities[e.Group] = group
	}

	return grEntities
}

// generate string in dot format from entities
func RenderViewToDotFormat(entities *[]Entity, appendDot string) string {
	output := bytes.NewBufferString("digraph G { \n")

	if appendDot != `` {
		output.WriteString(`  `) // padding
		output.WriteString(appendDot)
		output.WriteString("\n")
	}

	sort.Sort(byGroup(*entities)) // sort by groups and names
	grEntities := groupEntitiesArr(entities)

	for _, group := range grEntities {
		group.Wrap(output)
	}

	output.WriteString(`}`)

	return output.String()
}

func addGraphNodes(entities *[]Entity, output *bytes.Buffer, padding string) {
	for _, entity := range *entities {
		addNode(output, entity, padding)
		for _, child := range entity.Children {
			output.WriteString(padding) // padding
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
}

// add node definition
func addNode(output *bytes.Buffer, entity Entity, padding string) {
	output.WriteString(padding) // padding
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
