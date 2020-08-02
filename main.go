package main

import (
	"net/http"
	"practicing/transport"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/add", transport.AddArticle).Methods("POST")
	r.HandleFunc("/articles/delete", transport.DeleteArticle).Methods("POST")
	r.HandleFunc("/articles/get_all", transport.GetArtcles).Methods("GET")
	r.HandleFunc("/content_market/add", transport.AddContentMarket).Methods("POST")
	r.HandleFunc("/content_market/delete", transport.DeleteContentMarket).Methods("POST")
	r.HandleFunc("/content_market/get_all", transport.AllContentMarketing).Methods("GET")

	http.ListenAndServe(":6666", r)
}
