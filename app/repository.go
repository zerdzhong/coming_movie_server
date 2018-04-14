package app

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "mongodb://mongo:27017"

// DBNAME the name of the DB instance
const DBNAME = "douban_film_db"

// FILMCOLNAME the name of the document
const FILMCOLNAME = "film"

// COMINGFILMCOLNAME the name of the document
const COMINGFILMCOLNAME = "coming_films"

// GetComingFilms 获取即将上映电影
func (r Repository) GetComingFilms() DoubanComingFilms {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()
	c := session.DB(DBNAME).C(COMINGFILMCOLNAME)
	comingFilms := DoubanComingFilms{}
	if err := c.Find(nil).All(&comingFilms); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return comingFilms
}

// GetFilm 获取电影信息
func (r Repository) GetFilm(filmID string) DoubanFilms {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()
	c := session.DB(DBNAME).C(FILMCOLNAME)
	films := DoubanFilms{}
	if err := c.Find(bson.M{"id": filmID}).All(&films); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return films
}
