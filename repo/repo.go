package repo

import (
	"context"
	"database/sql"
	"fmt"
	"practicing/repo/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "postgres"
)

func AddArcicleDB(ctx context.Context, arcticle postgres.AddArticleParams) (postgres.Article, error) {
	//TODO: make pointers
	db, err := DBInit()
	//check if it's wont have a problem
	art, err := db.AddArticle(ctx, arcticle)
	if err != nil {
		return postgres.Article{}, err
	}
	return art, nil
}

func GetArticlesDB(ctx context.Context) ([]postgres.Article, error) {
	db, err := DBInit()
	if err != nil {
		return nil, err
	}
	art, err := db.AllArticles(ctx)
	if err != nil {
		return nil, err
	}
	return art, nil
}

func ArticleDeletDB(ctx context.Context, harvestID string) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	err = db.DeleteArticles(ctx, harvestID)
	if err != nil {
		return err
	}
	return nil
}

func AddContentMarketDB(ctx context.Context, contMarket postgres.AddContentMarketingParams) (postgres.Contentmarketing, error) {
	db, err := DBInit()
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	content, err := db.AddContentMarketing(ctx, contMarket)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	return content, nil
}

func DeleteContentMarketDB(ctx context.Context, harvestID string) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	err = db.DeleteContentMarketing(ctx, harvestID)
	if err != nil {
		return err
	}
	return nil
}

func AllContentMarketingDB(ctx context.Context) ([]postgres.Contentmarketing, error) {
	db, err := DBInit()
	if err != nil {
		return nil, err
	}
	answer, err := db.AllContentMarketing(ctx)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func DBInit() (*postgres.Queries, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	db := postgres.New(conn)
	return db, nil
}
