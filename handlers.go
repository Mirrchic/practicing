package main

import (
	"encoding/json"
	"net/http"
)

func AddArticle(w http.ResponseWriter, r *http.Request) {
	answer, err := AddArticleFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	err := ArcicleDeleteFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Secsessufully deleted"))
}

func GetArtcles(w http.ResponseWriter, r *http.Request) {
	answer, err := GetArticlesFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func AddContentMarket(w http.ResponseWriter, r *http.Request) {
	answer, err := AddContentMarkeFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func AllContentMarketing(w http.ResponseWriter, r *http.Request) {
	answer, err := AllContentMarketingFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func DeleteContentMarket(w http.ResponseWriter, r *http.Request) {
	err := DeleteContentMarketFetch(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Sucsessufully deleted"))
}
