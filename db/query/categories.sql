-- name: CreateCategories :one

INSERT INTO categories (
    id,
    name,
    type
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;