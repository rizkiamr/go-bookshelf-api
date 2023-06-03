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
    "insertedAt",
    "updatedAt"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
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
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;