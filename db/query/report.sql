-- name: GetReportbyUser :one
SELECT user_id, 
       SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) AS total_income,
       SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) AS total_expense
FROM transaction
WHERE user_id = $1
GROUP BY user_id;

-- name: GetDetailsReportbyUser :many
SELECT t.id, t.user_id, t.category_id, c.name AS category_name, t.amount, t.description, t.created_at
FROM transaction t
JOIN categories c ON t.category_id = c.id
WHERE t.user_id = $1
ORDER BY t.created_at DESC;

-- name: GetReportbyDate :many
SELECT t.id, t.user_id, t.category_id, c.name AS category_name, t.amount, t.type, t.description, t.created_at
FROM transaction t
JOIN categories c ON t.category_id = c.id
WHERE t.user_id = $1
AND t.created_at BETWEEN $2 AND $3
ORDER BY t.created_at DESC;

-- name: GetReportbyCategory :many
SELECT c.name AS category_name,
       SUM(CASE WHEN t.type = 'income' THEN t.amount ELSE 0 END) AS total_income,
       SUM(CASE WHEN t.type = 'expense' THEN t.amount ELSE 0 END) AS total_expense
FROM transaction t
JOIN categories c ON t.category_id = c.id
GROUP BY c.name;


-- name: CreateReportUser :exec
INSERT INTO reports (
  id,
  user_id,
  period,
  total_income,
  total_expense
)
SELECT
  $2, -- ID passed as a parameter
  t.user_id,
  'Daily',
  SUM(CASE WHEN t.type = 'income' THEN t.amount ELSE 0 END),
  SUM(CASE WHEN t.type = 'expense' THEN t.amount ELSE 0 END)
FROM transaction t
WHERE t.user_id = $1
AND t.created_at BETWEEN NOW() - INTERVAL '1 day' AND NOW()
GROUP BY t.user_id;
