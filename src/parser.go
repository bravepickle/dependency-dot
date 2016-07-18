// CSV parser

package main

import "encoding/csv"

import "io/ioutil"

import "io"
import "os"

import "fmt"
import "strings"

var Debug bool

const dbgLimit = 100

type Entity struct {
	Id     string
	Name   string
	RefIds []string

	Children   []*Entity
	ParentNode *Entity
}

func (e *Entity) RootNode() *Entity {
	if e.ParentNode != nil {
		return e.ParentNode.RootNode()
	} else {
		return nil
	}
}

func ParseCsv(filename string, columnsToParse map[string]string) {
	//	data, err := ioutil.ReadFile(filename)
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	if Debug {
		//		parsedStr := string(data)
		parsedStr, _ := ioutil.ReadAll(file)
		if len(parsedStr) > dbgLimit {
			fmt.Println(`Found file contents:`, "\n", string(parsedStr[:dbgLimit]), `...`)
		} else {
			fmt.Println(`Found file contents:`, "\n", string(parsedStr))
		}

		file.Seek(0, os.SEEK_SET) // rewind to start of file
	}

	csvReader := csv.NewReader(file)
	//	var rows [][]string

	titles, err := csvReader.Read()
	if err == io.EOF {
		panic(`File is empty`)
	}

	colMap := map[string]int{}
	for colKey, colName := range columnsToParse {
		for i, s := range titles {
			if colName == s {
				colMap[colKey] = i

				break
			}
		}

		if _, ok := colMap[colKey]; !ok {
			panic(fmt.Sprintln(`Field not found within list of titles:`, colName))
		}
	}

	if Debug {
		fmt.Println(`Mapped titles and columns:`, colMap)
	}

	var entities []Entity
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(record[colMap[`id`]])
		//		if _, ok := record[colMap[`id`]]; ok {
		//			panic(fmt.Sprintln(`Field not found for record:`, record))
		//		}

		var entity = Entity{Id: record[colMap[`id`]], Name: record[colMap[`name`]], RefIds: strings.Split(record[colMap[`id`]], `,`)}

		fmt.Println(entity)

		entities = append(entities, entity)

		//		break
	}

	// TODO: add children and root nodes. Ensure that infinite loop is not possible. Fix bug with "./parser.go:95: assignment count mismatch: 2 = 1"
	//		rows, err := csvReader.ReadAll()

	//	if err != nil {
	//		panic(err)
	//	}

	//	fmt.Println(string(rows))
}
