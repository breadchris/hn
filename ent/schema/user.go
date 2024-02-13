package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().Comment("The user's unique username. Case-sensitive."),
		field.Int64("created").Optional().Comment("Creation date of the user, in Unix Time."),
		field.Int("karma").Default(0).Comment("The user's karma."),
		field.Text("about").Optional().Nillable().Comment("The user's optional self-description. HTML."),
		// Assuming 'submitted' is a list of IDs referencing stories, polls, and comments
		// This field might be better represented as an edge to those entities, rather than a field
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	// Define edges to stories, polls, and comments here if needed
	return nil
}
