package main

//DoubanFilm film model
type DoubanFilm struct {
	// ID         bson.ObjectId `bson:"_id"`
	Name       string   `json:"name"`
	DoubanID   string   `json:"id"`
	Director   string   `json:"director"`
	Writer     string   `json:"writer"`
	Actors     []string `json:"actor"`
	PostImgURL string   `json:"post_img_url"`
	DetailURL  string   `json:"detail_url"`
	Rate       float32  `json:"rate"`
	RateSum    int32    `json:"rate_sum"`
}

//DoubanComingFilm 即将上映 model
type DoubanComingFilm struct {
	DetailURL      string `json:"detail_url"`
	PlayDate       string `json:"play_date"`
	Name           string `json:"name"`
	WantWatchCount string `json:"wish_watch_count"`
}

//DoubanFilms Array of DoubanFilm
type DoubanFilms []DoubanFilm
