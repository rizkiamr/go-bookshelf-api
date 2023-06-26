// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: author.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
    id,
    name
) VALUES (
    $1,
    $2
)
RETURNING id, name, "insertedAt"
`

type CreateAuthorParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.ID, arg.Name)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.InsertedAt)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

func (q *Queries) DeleteAuthor(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, "insertedAt" FROM authors
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id string) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.InsertedAt)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, "insertedAt" FROM authors
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAuthorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Author{}
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.InsertedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2
WHERE id = $1
RETURNING id, name, "insertedAt"
`

type UpdateAuthorParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor, arg.ID, arg.Name)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.InsertedAt)
	return i, err
}
