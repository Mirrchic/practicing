-- name: AllNews :many
SELECT * FROM news;

-- name: GetNews :one
SELECT * FROM news 
WHERE id = $1 LIMIT 1;

-- name: DeleteNews :exec
DELETE FROM news 
WHERE id = $1;

-- name: AddNews :many
INSERT INTO news (newstitle, newsstatus)
VALUES ($1, $2) RETURNING *;
