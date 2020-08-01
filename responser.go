package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"practicing/postgres"
)

func AddArticleFetch(r *http.Request) (postgres.Article, error) {
	//TODO: make pointers
	art, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return postgres.Article{}, err
	}
	var article postgres.AddArticleParams
	err = json.Unmarshal(art, &article)
	if err != nil {
		return postgres.Article{}, err
	}
	artReturn, err := AddArcicleDB(article)
	if err != nil {
		return postgres.Article{}, err
	}
	return artReturn, nil
	//todo
}

//todo Get Article
func GetArticlesFetch(r *http.Request) ([]postgres.Article, error) {
	ctx := r.Context()
	art, err := GetArticlesDB(ctx)
	if err != nil {
		return nil, err
	}
	return art, nil

}

func ArcicleDeleteFetch(r *http.Request) error {
	art, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var article *postgres.Article
	err = json.Unmarshal(art, &article)
	if err != nil {
		return err
	}
	err = ArticleDeletDB(article.Harvestid)
	if err != nil {
		return err
	}
	return nil
}
