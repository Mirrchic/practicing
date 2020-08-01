package main

import (
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
	r.HandleFunc("/articles/add", AddArticle).Methods("POST")
	r.HandleFunc("/articles/delete", DeleteArticle).Methods("POST")
	r.HandleFunc("/articles/get_all", GetArtcles).Methods("GET")
	r.HandleFunc("/content_market/add", AddContentMarket).Methods("POST")
	r.HandleFunc("/content_market/delete", DeleteContentMarket).Methods("POST")
	r.HandleFunc("/content_market/get_all", AllContentMarketing).Methods("GET")

	http.ListenAndServe(":6666", r)
}
