-- name: CreatePublisher :one
INSERT INTO publishers (
    name,
    "insertedAt"
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetPublisher :one
SELECT * FROM publishers
WHERE id = $1 LIMIT 1;

-- name: ListPublishers :many
SELECT * FROM publishers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePublisher :one
UPDATE publishers
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePublisher :exec
DELETE FROM publishers
WHERE id = $1;