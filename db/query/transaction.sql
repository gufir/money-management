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
