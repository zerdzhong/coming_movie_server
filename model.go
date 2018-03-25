package main

import "gopkg.in/mgo.v2/bson"

// DoubanFilm film model
type DoubanFilm struct {
	ID         bson.ObjectId `bson:"_id"`
	NAME       string        `json:"name"`
	DoubanID   string        `json:"id"`
	Director   string        `json:"director"`
	Writer     string        `json:"writer"`
	Actors     []string      `json:"actor"`
	PostImgURL string        `json:"post_img_url"`
	DetailURL  string        `json:"detail_url"`
	Rate       float32       `json:"rate"`
	RateSum    int32         `json:"rate_sum"`
}

//DoubanFilms Array of DoubanFilm
type DoubanFilms []DoubanFilm
