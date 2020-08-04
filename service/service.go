package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"practicing/repo"
	"practicing/repo/postgres"
)

type Service struct {
	Request *http.Request
}

func New(r *http.Request) *Service {
	return &Service{
		Request: r,
	}
}

func (r *Service) AddArticleFetch() (postgres.Article, error) {
	//TODO: make pointers
	art, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return postgres.Article{}, err
	}
	var article postgres.AddArticleParams
	err = json.Unmarshal(art, &article)
	if err != nil {
		return postgres.Article{}, err
	}
	artReturn, err := repo.AddArcicleDB(r.Request.Context(), article)
	if err != nil {
		return postgres.Article{}, err
	}
	return artReturn, nil
	//todo
}

//todo Get Article
func (r *Service) GetArticlesFetch() ([]postgres.Article, error) {
	ctx := r.Request.Context()
	art, err := repo.GetArticlesDB(ctx)
	if err != nil {
		return nil, err
	}
	return art, nil

}

func (r *Service) ArcicleDeleteFetch() error {
	art, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return err
	}
	var article *postgres.Article
	err = json.Unmarshal(art, &article)
	if err != nil {
		return err
	}
	err = repo.ArticleDeletDB(r.Request.Context(), article.Harvestid)
	if err != nil {
		return err
	}
	return nil
}

func (r *Service) AddContentMarkeFetch() (postgres.Contentmarketing, error) {
	contentJson, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	var market postgres.AddContentMarketingParams
	err = json.Unmarshal(contentJson, &market)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	answer, err := repo.AddContentMarketDB(r.Request.Context(), market)
	if err != nil {
		return postgres.Contentmarketing{}, err
	}
	return answer, nil
}

func (r *Service) DeleteContentMarketFetch() error {
	contentDelJson, err := ioutil.ReadAll(r.Request.Body)
	if err != nil {
		return err
	}
	var marketDelete postgres.Contentmarketing
	err = json.Unmarshal(contentDelJson, &marketDelete)
	if err != nil {
		return err
	}
	err = repo.DeleteContentMarketDB(r.Request.Context(), marketDelete.Harvestid)
	if err != nil {
		return err
	}
	return nil
}

func (r *Service) AllContentMarketingFetch() ([]postgres.Contentmarketing, error) {
	ctx := r.Request.Context()
	answer, err := repo.AllContentMarketingDB(ctx)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
