// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: posts.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  cate_id,
  user_id,
  title,
  content
) VALUES (
  $1, $2, $3, $4
) RETURNING id, cate_id, user_id, title, content, created_at
`

type CreatePostParams struct {
	CateID  sql.NullInt32 `json:"cate_id"`
	UserID  sql.NullInt32 `json:"user_id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.CateID,
		arg.UserID,
		arg.Title,
		arg.Content,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CateID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :one
SELECT id, cate_id, user_id, title, content, created_at FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPosts(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPosts, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CateID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, cate_id, user_id, title, content, created_at FROM posts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CateID,
			&i.UserID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
