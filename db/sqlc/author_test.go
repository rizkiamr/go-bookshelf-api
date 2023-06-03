package db

import (
	"context"
	"testing"
	"time"

	"github.com/rizkiamr/go-bookshelf-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomAuthor(t *testing.T) Author {
	name := util.RandomName()

	author, err := testQueries.CreateAuthor(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, name, author.Name)

	require.NotZero(t, author.ID)
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