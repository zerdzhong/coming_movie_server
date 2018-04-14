package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var repos = Repository{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	comingFilms := repos.GetComingFilms()
	data, _ := json.Marshal(comingFilms)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func comingHandler(w http.ResponseWriter, r *http.Request) {
	comingFilms := repos.GetComingFilms()
	data, _ := json.Marshal(comingFilms)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func filmHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmID := vars["filmId"]

	film := repos.GetFilm(filmID)
	data, _ := json.Marshal(film)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
