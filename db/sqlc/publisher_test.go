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

func createRandomPublisher(t *testing.T) Publisher {

	uuidObj := uuid.New()
	uuidStr := uuidObj.String()
	uuidWithoutDashes := strings.ReplaceAll(uuidStr, "-", "")

	arg := CreatePublisherParams{
		ID:   uuidWithoutDashes,
		Name: util.RandomName(),
	}

	publisher, err := testQueries.CreatePublisher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, publisher)

	require.Equal(t, arg.ID, publisher.ID)
	require.Equal(t, arg.Name, publisher.Name)

	require.NotZero(t, publisher.InsertedAt)

	return publisher
}

func TestCreatePublisher(t *testing.T) {
	createRandomPublisher(t)
}

func TestGetPublisher(t *testing.T) {
	publisher1 := createRandomPublisher(t)
	publisher2, err := testQueries.GetPublisher(context.Background(), publisher1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, publisher2)

	require.Equal(t, publisher1.ID, publisher2.ID)
	require.Equal(t, publisher1.Name, publisher2.Name)

	require.WithinDuration(t, publisher1.InsertedAt, publisher2.InsertedAt, time.Second)
}

func TestUpdatePublisher(t *testing.T) {
	publisher1 := createRandomPublisher(t)

	arg := UpdatePublisherParams{
		ID:   publisher1.ID,
		Name: util.RandomName(),
	}

	publisher2, err := testQueries.UpdatePublisher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, publisher2)

	require.Equal(t, publisher1.ID, publisher2.ID)
	require.Equal(t, arg.Name, publisher2.Name)
	require.WithinDuration(t, publisher1.InsertedAt, publisher2.InsertedAt, time.Second)
}

func TestDeletePublisher(t *testing.T) {
	publisher1 := createRandomPublisher(t)
	err := testQueries.DeletePublisher(context.Background(), publisher1.ID)
	require.NoError(t, err)

	publisher2, err := testQueries.GetPublisher(context.Background(), publisher1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, publisher2)
}

func TestListPublishers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPublisher(t)
	}

	arg := ListPublishersParams{
		Limit:  5,
		Offset: 5,
	}

	publishers, err := testQueries.ListPublishers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, publishers, 5)

	for _, publisher := range publishers {
		require.NotEmpty(t, publisher)
	}
}
