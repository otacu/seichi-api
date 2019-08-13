package main

type GetAnimeResponse struct {
	Id   uint   `json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
}

type GetSeichiResponse struct {
	Id      uint    `json:"id" form:"id" query:"id"`
	Name    string  `json:"name" form:"name" query:"name"`
	AnimeId uint    `json:"animeId" form:"animeId" query:"animeId"`
	Lng     float32 `json:"lng" form:"lng" query:"lng"`
	Lat     float32 `json:"lat" form:"lat" query:"lat"`
}
