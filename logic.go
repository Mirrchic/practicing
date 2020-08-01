package main

import (
	"context"
	"database/sql"
	"fmt"
	"practicing/postgres"
)

func AddArcicleDB(arcticle postgres.AddArticleParams) (postgres.Article, error) {
	//TODO: make pointers
	db, err := DBInit()
	//check if it's wont have a problem
	art, err := db.AddArticle(context.Background(), arcticle)
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

func ArticleDeletDB(harvestID string) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	err = db.DeleteArticles(context.Background(), harvestID)
	if err != nil {
		return err
	}
	return nil
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
