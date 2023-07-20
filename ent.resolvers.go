package bookcatalog

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/pratikjagrut/book-catalog/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*ent.Author, error) {
	return r.client.Author.Query().All(ctx)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*ent.Book, error) {
	return r.client.Book.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
