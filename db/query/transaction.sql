-- name: CreateTransaction :one
INSERT INTO "transaction" (
    id,
    user_id,
    amount,
    type,
    description,
    category_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) 
RETURNING *;

-- name: GetTransaction :one
SELECT * 
FROM "transaction"
WHERE id = $1;

-- name: GetTransactionByType :many
SELECT *
FROM "transaction"
WHERE type = $1 AND user_id = $2;

-- name: GetTransactionByuserId :many
SELECT t.id, description, amount, type, created_at, updated_at, deleted_at, name, category_id, user_id
FROM "transaction" t
JOIN "categories" c
ON t.category_id = c.id
WHERE user_id = $1;

-- name: UpdateTransaction :one
UPDATE "transaction"
SET
    amount = COALESCE(sqlc.narg(amount), amount),
    type = COALESCE(sqlc.narg(type), type),
    description = COALESCE(sqlc.narg(description), description),
    category_id = COALESCE(sqlc.arg(category_id), category_id)
WHERE
    user_id = sqlc.arg(user_id) AND id = sqlc.arg(id)
RETURNING *;
