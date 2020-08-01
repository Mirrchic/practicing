package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "postgres"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("articles/add_ar", AddArticle).Methods("POST")
	r.HandleFunc("articles/delete", DeleteArticle).Methods("POST")
	r.HandleFunc("articles/get_articles", GetArtcles).Methods("GET")
	http.ListenAndServe(":8888", nil)

}

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
