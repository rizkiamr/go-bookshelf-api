package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	name := "penulis1"

	author, err := testQueries.CreateAuthor(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, name, author.Name)

	require.NotZero(t, author.ID)
	require.NotZero(t, author.InsertedAt)
}
