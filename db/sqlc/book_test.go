package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/rizkiamr/go-bookshelf-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomBook(t *testing.T) Book {
	author1 := createRandomAuthor(t)
	publisher1 := createRandomPublisher(t)
	tz, _ := time.LoadLocation("Etc/UTC")

	arg := CreateBookParams{
		Name:        util.RandomName(),
		Year:        util.RandomInt32(1000, 2100),
		AuthorID:    author1.ID,
		Summary:     util.RandomString(15),
		PublisherID: publisher1.ID,
		PageCount:   util.RandomInt32(0, 1000),
		ReadPage:    util.RandomInt32(0, 1000),
		Finished:    util.RandomBool(),
		Reading:     util.RandomBool(),
		UpdatedAt:   sql.NullTime{Time: time.Now().In(tz), Valid: true},
	}

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.NotZero(t, book.ID)
	require.NotZero(t, book.Name)
	require.NotZero(t, book.Year)
	require.NotZero(t, book.AuthorID)
	require.NotZero(t, book.Summary)
	require.NotZero(t, book.PublisherID)
	require.NotZero(t, book.PageCount)
	require.NotZero(t, book.ReadPage)
	require.NotNil(t, book.Finished)
	require.NotNil(t, book.Reading)
	require.NotZero(t, book.InsertedAt)
	require.NotZero(t, book.UpdatedAt)

	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book1 := createRandomBook(t)
	book2, err := testQueries.GetBook(context.Background(), book1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book1.Name, book2.Name)
	require.Equal(t, book1.Year, book2.Year)
	require.Equal(t, book1.AuthorID, book2.AuthorID)
	require.Equal(t, book1.Summary, book2.Summary)
	require.Equal(t, book1.PublisherID, book2.PublisherID)
	require.Equal(t, book1.PageCount, book2.PageCount)
	require.Equal(t, book1.ReadPage, book2.ReadPage)
	require.Equal(t, book1.Finished, book2.Finished)
	require.Equal(t, book1.Reading, book2.Reading)
	require.Equal(t, book1.UpdatedAt, book2.UpdatedAt)

	require.WithinDuration(t, book1.InsertedAt, book2.InsertedAt, time.Second)
}

func TestUpdateBook(t *testing.T) {
	book1 := createRandomBook(t)
	author1 := createRandomAuthor(t)
	publisher1 := createRandomPublisher(t)

	arg := UpdateBookParams{
		ID:   book1.ID,
		Name: util.RandomName(),
		Year: util.RandomInt32(1, 3000),
		AuthorID: author1.ID,
		Summary: util.RandomString(15),
		PublisherID: publisher1.ID,
		PageCount: util.RandomInt32(1, 5000),
		ReadPage: util.RandomInt32(1, 100),
		Finished: util.RandomBool(),
		Reading: util.RandomBool(),
		UpdatedAt: book1.UpdatedAt,
	}

	book2, err := testQueries.UpdateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, arg.Name, book2.Name)
	require.Equal(t, arg.Year, book2.Year)
	require.Equal(t, arg.AuthorID, book2.AuthorID)
	require.Equal(t, arg.Summary, book2.Summary)
	require.Equal(t, arg.PublisherID, book2.PublisherID)
	require.Equal(t, arg.PageCount, book2.PageCount)
	require.Equal(t, arg.ReadPage, book2.ReadPage)
	require.Equal(t, arg.Finished, book2.Finished)
	require.Equal(t, arg.Reading, book2.Reading)
	require.Equal(t, book1.UpdatedAt, book2.UpdatedAt)

	require.WithinDuration(t, book1.InsertedAt, book2.InsertedAt, time.Second)
}

func TestDeleteBook(t *testing.T) {
	book1 := createRandomBook(t)
	err := testQueries.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)

	book2, err := testQueries.GetBook(context.Background(), book1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, book2)
}

func TestListBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBook(t)
	}

	arg := ListBooksParams{
		Limit:  5,
		Offset: 5,
	}

	books, err := testQueries.ListBooks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, books, 5)

	for _, book := range books {
		require.NotEmpty(t, book)
	}
}
