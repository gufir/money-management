-- name: CreateVerifyEmails :one
INSERT INTO verify_emails (
    user_id,
    email,
    secret_code
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: UpdateVerifyEmails :one
UPDATE verify_emails
SET
    is_used = true
WHERE
    id = @id
    AND secret_code = @secret_code
    AND is_used = false
    AND expires_at > now()
RETURNING *;