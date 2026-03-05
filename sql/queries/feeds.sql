-- name: Addfeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: Getfeeds :many
Select * from feeds;

-- name: Getfeedurl :one
Select * from feeds where url = $1;