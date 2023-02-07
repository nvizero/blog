-- name: CreatePost :one
INSERT INTO posts (
  cate_id,
  user_id,
  title,
  content
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPosts :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;
