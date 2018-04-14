package app

import "gopkg.in/mgo.v2/bson"

//DoubanFilm film model
type DoubanFilm struct {
	// ID         bson.ObjectId `bson:"_id"`
	Name       []string `bson:"name"`
	DoubanID   string   `bson:"id"`
	Director   []string `bson:"director"`
	Writer     []string `bson:"writer"`
	Actors     []string `bson:"actor"`
	PostImgURL []string `bson:"post_img_url"`
	DetailURL  string   `bson:"detail_url"`
	Rate       float32  `bson:"rate"`
	RateSum    int32    `bson:"rate_sum"`
}

//DoubanComingFilm 即将上映 model
type DoubanComingFilm struct {
	ID             bson.ObjectId `bson:"_id"`
	DetailURL      string        `bson:"detail_url"`
	PlayDate       string        `bson:"play_date"`
	Name           string        `bson:"name"`
	WantWatchCount string        `bson:"wish_watch_count"`
}

//DoubanFilms Array of DoubanFilm
type DoubanFilms []DoubanFilm

//DoubanComingFilms Array of DoubanComingFilm
type DoubanComingFilms []DoubanComingFilm
