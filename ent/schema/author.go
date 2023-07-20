package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}

func (Author) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		// entgql.RelayConnection(),
	}
}
