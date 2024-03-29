// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/breadchris/hn/ent/item"
)

// CreateItemInput represents a mutation input for creating items.
type CreateItemInput struct {
	Deleted     bool
	Type        item.Type
	By          *string
	Text        *string
	Dead        bool
	Parent      *int
	Poll        *int
	Kids        []int
	URL         *string
	Score       *int
	Title       *string
	Parts       []int
	Descendants *int
	Time        int
	ChildIDs    []string
	ParentIDs   []string
}

// Mutate applies the CreateItemInput on the ItemMutation builder.
func (i *CreateItemInput) Mutate(m *ItemMutation) {
	m.SetDeleted(i.Deleted)
	m.SetType(i.Type)
	if v := i.By; v != nil {
		m.SetBy(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	m.SetDead(i.Dead)
	if v := i.Parent; v != nil {
		m.SetParent(*v)
	}
	if v := i.Poll; v != nil {
		m.SetPoll(*v)
	}
	if v := i.Kids; v != nil {
		m.SetKids(v)
	}
	if v := i.URL; v != nil {
		m.SetURL(*v)
	}
	if v := i.Score; v != nil {
		m.SetScore(*v)
	}
	if v := i.Title; v != nil {
		m.SetTitle(*v)
	}
	if v := i.Parts; v != nil {
		m.SetParts(v)
	}
	if v := i.Descendants; v != nil {
		m.SetDescendants(*v)
	}
	m.SetTime(i.Time)
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.ParentIDs; len(v) > 0 {
		m.AddParentIDs(v...)
	}
}

// SetInput applies the change-set in the CreateItemInput on the ItemCreate builder.
func (c *ItemCreate) SetInput(i CreateItemInput) *ItemCreate {
	i.Mutate(c.Mutation())
	return c
}
