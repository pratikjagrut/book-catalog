package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("genre").NotEmpty(),
		field.String("publication_date").NotEmpty(),
		field.String("isbn").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", Author.Type).
			Ref("books").
			Unique(),
	}
}

func (Book) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
		// entgql.RelayConnection(),
	}
}
