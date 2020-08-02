package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"practicing/repo"
	"practicing/repo/postgres"
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
	artReturn, err := repo.AddArcicleDB(article)
	if err != nil {
		return postgres.Article{}, err
	}
	return artReturn, nil
	//todo
}

//todo Get Article
func GetArticlesFetch(r *http.Request) ([]postgres.Article, error) {
	ctx := r.Context()
	art, err := repo.GetArticlesDB(ctx)
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
	err = repo.ArticleDeletDB(article.Harvestid)
	if err != nil {
		return err
	}
	return nil
}

func AddContentMarkeFetch(r *http.Request) (postgres.Contentmarketing, error) {
	contentJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	var market postgres.AddContentMarketingParams
	err = json.Unmarshal(contentJson, &market)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	answer, err := repo.AddContentMarketDB(market)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	return answer, nil
}

func DeleteContentMarketFetch(r *http.Request) error {
	contentDelJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	var marketDelete postgres.Contentmarketing
	err = json.Unmarshal(contentDelJson, &marketDelete)
	if err != nil {
		return err
	}
	err = repo.DeleteContentMarketDB(marketDelete.Harvestid)
	if err != nil {
		return err
	}
	return nil
}

func AllContentMarketingFetch(r *http.Request) ([]postgres.Contentmarketing, error) {
	ctx := r.Context()
	answer, err := repo.AllContentMarketingDB(ctx)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
