-- name: CreateCate :one
INSERT INTO cates (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetCate :one
SELECT * FROM cates
WHERE id = $1 LIMIT 1;

-- name: ListCates :many
SELECT * FROM cates
ORDER BY id
LIMIT $1
OFFSET $2;
