// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package postgres

import (
	"context"
)

const addArticle = `-- name: AddArticle :one
INSERT INTO articles (harvestId, cerebroScore, url, title, cleanImage )
VALUES ($1, $2, $3, $4, $5) RETURNING harvestid, cerebroscore, url, title, cleanimage
`

type AddArticleParams struct {
	Harvestid    string `json:"harvestid"`
	Cerebroscore string `json:"cerebroscore"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	Cleanimage   string `json:"cleanimage"`
}

func (q *Queries) AddArticle(ctx context.Context, arg AddArticleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, addArticle,
		arg.Harvestid,
		arg.Cerebroscore,
		arg.Url,
		arg.Title,
		arg.Cleanimage,
	)
	var i Article
	err := row.Scan(
		&i.Harvestid,
		&i.Cerebroscore,
		&i.Url,
		&i.Title,
		&i.Cleanimage,
	)
	return i, err
}

const addContentMarketing = `-- name: AddContentMarketing :one
INSERT INTO contentMarketing (harvestId, commercialPartner, logoURL, cerebroScore)
VALUES ($1, $2, $3, $4) RETURNING harvestid, commercialpartner, logourl, cerebroscore
`

type AddContentMarketingParams struct {
	Harvestid         string `json:"harvestid"`
	Commercialpartner string `json:"commercialpartner"`
	Logourl           string `json:"logourl"`
	Cerebroscore      string `json:"cerebroscore"`
}

func (q *Queries) AddContentMarketing(ctx context.Context, arg AddContentMarketingParams) (Contentmarketing, error) {
	row := q.db.QueryRowContext(ctx, addContentMarketing,
		arg.Harvestid,
		arg.Commercialpartner,
		arg.Logourl,
		arg.Cerebroscore,
	)
	var i Contentmarketing
	err := row.Scan(
		&i.Harvestid,
		&i.Commercialpartner,
		&i.Logourl,
		&i.Cerebroscore,
	)
	return i, err
}

const allArticles = `-- name: AllArticles :many
SELECT harvestid, cerebroscore, url, title, cleanimage FROM articles
`

func (q *Queries) AllArticles(ctx context.Context) ([]Article, error) {
	rows, err := q.db.QueryContext(ctx, allArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.Harvestid,
			&i.Cerebroscore,
			&i.Url,
			&i.Title,
			&i.Cleanimage,
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

const allContentMarketing = `-- name: AllContentMarketing :many
SELECT harvestid, commercialpartner, logourl, cerebroscore FROM contentMarketing
`

func (q *Queries) AllContentMarketing(ctx context.Context) ([]Contentmarketing, error) {
	rows, err := q.db.QueryContext(ctx, allContentMarketing)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contentmarketing
	for rows.Next() {
		var i Contentmarketing
		if err := rows.Scan(
			&i.Harvestid,
			&i.Commercialpartner,
			&i.Logourl,
			&i.Cerebroscore,
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

const deleteArticles = `-- name: DeleteArticles :exec
DELETE FROM articles 
WHERE harvestId = $1
`

func (q *Queries) DeleteArticles(ctx context.Context, harvestid string) error {
	_, err := q.db.ExecContext(ctx, deleteArticles, harvestid)
	return err
}

const deleteContentMarketing = `-- name: DeleteContentMarketing :exec
DELETE FROM contentMarketing 
WHERE harvestId = $1
`

func (q *Queries) DeleteContentMarketing(ctx context.Context, harvestid string) error {
	_, err := q.db.ExecContext(ctx, deleteContentMarketing, harvestid)
	return err
}

const getNewsArticles = `-- name: GetNewsArticles :one
SELECT harvestid, cerebroscore, url, title, cleanimage FROM articles 
WHERE harvestId = $1 LIMIT 1
`

func (q *Queries) GetNewsArticles(ctx context.Context, harvestid string) (Article, error) {
	row := q.db.QueryRowContext(ctx, getNewsArticles, harvestid)
	var i Article
	err := row.Scan(
		&i.Harvestid,
		&i.Cerebroscore,
		&i.Url,
		&i.Title,
		&i.Cleanimage,
	)
	return i, err
}

const getNewsContentMarketing = `-- name: GetNewsContentMarketing :one
SELECT harvestid, cerebroscore, url, title, cleanimage FROM articles 
WHERE harvestId = $1 LIMIT 1
`

func (q *Queries) GetNewsContentMarketing(ctx context.Context, harvestid string) (Article, error) {
	row := q.db.QueryRowContext(ctx, getNewsContentMarketing, harvestid)
	var i Article
	err := row.Scan(
		&i.Harvestid,
		&i.Cerebroscore,
		&i.Url,
		&i.Title,
		&i.Cleanimage,
	)
	return i, err
}
