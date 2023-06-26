package db

import (
	"context"
	"database/sql"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rizkiamr/go-bookshelf-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Author {
	uuidObj := uuid.New()
	uuidStr := uuidObj.String()
	uuidWithoutDashes := strings.ReplaceAll(uuidStr, "-", "")

	arg := CreateAuthorParams{
		ID:   uuidWithoutDashes,
		Name: util.RandomName(),
	}

	author, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.ID, author.ID)
	require.Equal(t, arg.Name, author.Name)

	require.NotZero(t, author.InsertedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestGetAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)
	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, author1.Name, author2.Name)

	require.WithinDuration(t, author1.InsertedAt, author2.InsertedAt, time.Second)
}

func TestUpdateAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)

	arg := UpdateAuthorParams{
		ID:   author1.ID,
		Name: util.RandomName(),
	}

	author2, err := testQueries.UpdateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, arg.Name, author2.Name)
	require.WithinDuration(t, author1.InsertedAt, author2.InsertedAt, time.Second)
}

func TestDeleteAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)
	err := testQueries.DeleteAuthor(context.Background(), author1.ID)
	require.NoError(t, err)

	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, author2)
}

func TestListAuthors(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAuthor(t)
	}

	arg := ListAuthorsParams{
		Limit:  5,
		Offset: 5,
	}

	authors, err := testQueries.ListAuthors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, authors, 5)

	for _, author := range authors {
		require.NotEmpty(t, author)
	}
}
