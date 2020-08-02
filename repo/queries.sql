-- name: AllArticles :many
SELECT * FROM articles;

-- name: AllContentMarketing :many
SELECT * FROM contentMarketing;

-- name: GetNewsArticles :one
SELECT * FROM articles 
WHERE harvestId = $1 LIMIT 1;

-- name: GetNewsContentMarketing :one
SELECT * FROM articles 
WHERE harvestId = $1 LIMIT 1;

-- name: DeleteArticles :exec
DELETE FROM articles 
WHERE harvestId = $1;

-- name: DeleteContentMarketing :exec
DELETE FROM contentMarketing 
WHERE harvestId = $1;

-- name: AddArticle :one
INSERT INTO articles (harvestId, cerebroScore, url, title, cleanImage )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: AddContentMarketing :one
INSERT INTO contentMarketing (harvestId, commercialPartner, logoURL, cerebroScore)
VALUES ($1, $2, $3, $4) RETURNING *;
