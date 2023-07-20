package ent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestBookCatalog(t *testing.T) {
	// client, err := Open("sqlite3", "file:book-catalog.db?cache=shared&_fk=1")
	client, err := Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	assert.NoErrorf(t, err, "failed opening connection to sqlite")
	defer client.Close()

	ctx := context.Background()

	// Run the automatic migration tool to create all schema resources.
	err = client.Schema.Create(ctx)
	assert.NoErrorf(t, err, "failed creating schema resources")

	author, err := client.Author.Create().
		SetName("J. K. Rowling").
		SetEmail("jk@gmail.com").
		Save(ctx)
	assert.NoError(t, err)

	_, err = client.Book.Create().
		SetTitle("The Ink Black Heart").
		SetGenre("Mystery").
		SetIsbn("9780316413138").
		SetPublicationDate("30 August 2022").
		SetAuthor(author).
		Save(ctx)
	assert.NoError(t, err)

	author, err = client.Author.Create().
		SetName("George R. R. Martin").
		SetEmail("grrm@gmail.com").
		Save(ctx)
	assert.NoError(t, err)

	_, err = client.Book.Create().
		SetTitle("A Game of Thrones").
		SetGenre("Fantasy Fiction").
		SetIsbn("9780553593716").
		SetPublicationDate("1 August 1996").
		SetAuthor(author).
		Save(ctx)
	assert.NoError(t, err)

	books, err := client.Book.Query().All(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(books), 2)
}
