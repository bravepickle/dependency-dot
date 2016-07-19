// Entity model per each parsed record

package main

type Ref struct {
	Id    string // reference id
	Style string // reference style
}

type Entity struct {
	Id     string   // node id. Consists from EntityId and ReferencedId concatenated by `-`
	Name   string   // node name
	Style  string   // node style
	RefIds []string // subnodes' ids
	Refs   []Ref    // subnodes' references

	Children   []*Entity // subnodes
	ParentNode *Entity
}

// Get Ref by index
func (e *Entity) Ref(index int) Ref {
	//	if _, ok := colMap[colKey]; !ok {

	ref := e.Refs[index]
	//	if ref, ok := e.Refs[index]; ok {
	//	if ref {
	return ref
	//	} else {
	//		return nil
	//	}
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
		entity.ParentNode = e

		return true
	} else {
		return false
	}
}

func (e *Entity) RootNode() *Entity {
	if e.ParentNode != nil {
		return e.ParentNode.RootNode()
	} else {
		return nil
	}
}
