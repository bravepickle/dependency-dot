// Entity model per each parsed record

package main

import "strings"

type Ref struct {
	Id    string // current entity id
	RefId string // referenced entity id
	Style string // reference style
}

type Entity struct {
	Id     string   // node id. Consists from EntityId and ReferencedId concatenated by `-`
	Name   string   // node name
	Style  string   // node style
	Group  string   // node group it belongs to
	RefIds []string // subnodes' ids
	Refs   []Ref    // subnodes' references

	Children []*Entity // subnodes
}

// Get Ref by Referenced Entity ID. If not found, then second param will be false
func (e *Entity) Ref(id string) (Ref, bool) {
	for _, ref := range e.Refs {
		if ref.RefId == id {
			return ref, true
		}
	}

	return Ref{}, true
}

// check if child is added to list
func (e *Entity) HasChild(entity *Entity) bool {
	for _, child := range e.Children {
		if child.Id == entity.Id {
			return true
		}
	}

	return false
}

func (e *Entity) AddChild(entity *Entity) bool {
	if !e.HasChild(entity) {
		e.Children = append(e.Children, entity)

		return true
	} else {
		return false
	}
}

type byGroup []Entity

func (s byGroup) Less(i, j int) bool {
	if s[i].Group == s[j].Group {
		return strings.Compare(s[i].Name, s[j].Name) >= 0
	}

	return strings.Compare(s[i].Group, s[j].Group) >= 0
}

func (s byGroup) Len() int {
	return len(s)
}

func (s byGroup) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
