package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"practicing/postgres"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "users"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("while open \n", err)
	}

	db := postgres.New(conn)

	news, err := db.AddNews(context.Background(), postgres.AddNewsParams{
		Newsstatus: "Breaking",
		Newstitle:  "Somethink doesn't work sqlc",
	})
	if err != nil {
		log.Fatal("while created \n", err)
	}
	log.Print(news)
}
