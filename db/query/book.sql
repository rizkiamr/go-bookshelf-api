-- name: CreateBook :one
INSERT INTO books (
    name,
    year,
    author_id,
    summary,
    publisher_id,
    "pageCount",
    "readPage",
    finished,
    reading,
    "updatedAt"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBook :one
UPDATE books
SET
    name = $2,
    year = $3,
    author_id = $4,
    summary = $5,
    publisher_id = $6,
    "pageCount" = $7,
    "readPage" = $8,
    finished = $9,
    reading = $10,
    "updatedAt" = $11
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;