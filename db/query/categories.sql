-- name: CreateCategories :one

INSERT INTO categories (
    id,
    name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetCategories :many
SELECT *
FROM categories;
