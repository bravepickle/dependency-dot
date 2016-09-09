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

func ParseCsv(filename string, columnsToParse map[string]string) []Entity {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	if Debug {
		parsedStr, _ := ioutil.ReadAll(file)
		if len(parsedStr) > dbgLimit {
			fmt.Printf("Found file contents:\n%s ...\n\n", string(parsedStr[:dbgLimit]))
		} else {
			fmt.Printf("Found file contents:\n%s\n\n", string(parsedStr))
		}

		file.Seek(0, os.SEEK_SET) // rewind to start of file
	}

	csvReader := csv.NewReader(file)
	colMap := initColumnMap(&columnsToParse, csvReader)

	entities := initEntitiesFromCsv(csvReader, colMap)
	if Debug {
		fmt.Print(`Parsed entities: `, entities, "\n\n")
	}

	return entities
}

// init mapping of column titles and its index numbers
// function will search for titles of columns in the first row and move pointer to second line
func initColumnMap(columnsToParse *map[string]string, csvReader *csv.Reader) map[string]int {
	titles, err := csvReader.Read()
	if err == io.EOF {
		panic(`File is empty`)
	}

	colMap := map[string]int{}
	for colKey, colName := range *columnsToParse {
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

	return colMap
}

func parseRefIds(rawValue string, entityId string) (refIds []string, refs []Ref) {
	refIds = strings.Split(rawValue, `,`)
	for i, v := range refIds {
		refId := strings.TrimSpace(v)

		if refId == `` { // skip empty refs
			continue
		}

		refPieces := strings.Split(refId, `[`)
		refId = strings.TrimSpace(refPieces[0])

		ref := Ref{Id: entityId + `-` + refId, RefId: refId}

		if len(refPieces) > 1 {
			ref.Style = strings.TrimRight(refPieces[1], `]`)
		}

		refs = append(refs, ref)
		refIds[i] = refId
	}

	return refIds, refs
}

func initEntitiesFromCsv(csvReader *csv.Reader, colMap map[string]int) []Entity {
	var entities []Entity

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		refIds, refs := parseRefIds(record[colMap[`ref`]], record[colMap[`id`]])
		var entity = Entity{Id: record[colMap[`id`]], Name: record[colMap[`name`]], RefIds: refIds, Refs: refs, Style: record[colMap[`style`]]}
		entity.Group = strings.TrimSpace(record[colMap[`group`]])

		if entity.Id != `` && entity.Name != `` {
			entities = append(entities, entity)
		} else if Debug {
			fmt.Println(`Skipping entity with empty id or name:`, entity)
		}
	}

	addEntityDeps(&entities)

	return entities
}

// add entity dependencies between each other
func addEntityDeps(entities *[]Entity) {
	for i, entity := range *entities {
		for _, refId := range entity.RefIds {
			if refId == `` { // skip empty refIds
				continue
			}

			found := false
			for _, sub := range *entities {
				if sub.Id == refId {
					(*entities)[i].AddChild(&sub) // entity cannot be set directly

					found = true
					break
				}
			}

			if !found {
				fmt.Fprintf(os.Stderr, `Failed to find reference ID "%s" for row ID "%s"%s`, refId, entity.Id, "\n")
			}
		}
	}
}
