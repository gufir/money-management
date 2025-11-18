-- name: CreatePasswordReset :one
INSERT INTO password_reset_tokens (
    id,
    user_id,
    token,
    expired_at,
    is_used,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    FALSE,
    now()
) RETURNING *;

-- name: GetPasswordResetToken :one
SELECT * FROM password_reset_tokens
WHERE token = $1
LIMIT 1;

-- name: UsePasswordResetToken :exec
UPDATE password_reset_tokens
SET is_used = TRUE,
    used_at = now()
WHERE id = $1;