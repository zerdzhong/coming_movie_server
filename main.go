package main

import (
	"log"
	"net/http"

	"github.com/zerdzhong/coming_movie_server/app"
)

func main() {
	router := app.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
