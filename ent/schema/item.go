package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique(),
		field.Bool("deleted"),
		field.Enum("type").
			Values("job", "story", "comment", "poll", "pollopt"),
		field.String("by").
			Optional(),
		field.Text("text").
			Optional().
			Nillable(),
		field.Bool("dead"),
		field.Int("parent").Optional().Nillable(),
		field.Int("poll").Optional().Nillable(),
		field.JSON("kids", []int{}).Optional(),
		field.String("url").Optional().Nillable(),
		field.Int("score").Optional(),
		field.String("title").Optional().Nillable(),
		field.JSON("parts", []int{}).Optional(),
		field.Int("descendants").Optional(),
		field.Int("time").
			Annotations(
				entgql.OrderField("TIME"),
			),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Item.Type),
		edge.From("parents", Item.Type).
			Ref("children"),
	}
}

func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
