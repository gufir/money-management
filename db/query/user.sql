-- name: CreateUser :one

INSERT INTO users (
    username,
    full_name,
    email,
    hashed_password,
    user_uuid   
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users 
WHERE id = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email)
WHERE
    username = sqlc.arg(username)
RETURNING *;