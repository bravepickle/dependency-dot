// Entity model per each parsed record

package main

type Entity struct {
	Id     string
	Name   string
	RefIds []string

	Children   []*Entity
	ParentNode *Entity
}

func (e *Entity) HasChild(entity *Entity) bool {
	for _, refId := range e.RefIds {
		if refId == entity.Id {
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
