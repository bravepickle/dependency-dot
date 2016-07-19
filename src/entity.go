// Entity model per each parsed record

package main

type Ref struct {
	Id    string // current entity id
	RefId string // referenced entity id
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

// Get Ref by Referenced Entity ID. If not found, then second param will be false
func (e *Entity) Ref(id string) (Ref, bool) {
	//	if _, ok := colMap[colKey]; !ok {
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
