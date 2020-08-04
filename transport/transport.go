package transport

import (
	"encoding/json"
	"net/http"
	"practicing/service"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	answer, err := service.New(r).AddArticleFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	err := service.New(r).ArcicleDeleteFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Secsessufully deleted"))
}

func GetArtcles(w http.ResponseWriter, r *http.Request) {
	answer, err := service.New(r).GetArticlesFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func AddContentMarket(w http.ResponseWriter, r *http.Request) {
	answer, err := service.New(r).AddContentMarkeFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func AllContentMarketing(w http.ResponseWriter, r *http.Request) {
	answer, err := service.New(r).AllContentMarketingFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func DeleteContentMarket(w http.ResponseWriter, r *http.Request) {
	err := service.New(r).DeleteContentMarketFetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Sucsessufully deleted"))
}
