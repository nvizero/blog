-- name: CreatePost :one
INSERT INTO posts (
  title,
  cate_id,
  content
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;
