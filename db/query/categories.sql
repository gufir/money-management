-- name: CreateCategories :one

INSERT INTO categories (
    id,
    name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetAllCategories :many
SELECT *
FROM categories
ORDER BY name;

-- name: GetCategoryByName :one
SELECT * FROM categories 
WHERE name = $1 
LIMIT 1;

-- name: UpdateCategories :one
UPDATE categories
SET
    name = COALESCE(sqlc.narg(name), name)
WHERE 
    id = sqlc.arg(id)
RETURNING *;

-- name: GetCategoryById :one
SELECT * FROM categories
WHERE id = $1;