package main

// query是指get请求url中的参数
type GetSeichiRequest struct {
	AnimeId uint32 `json:"animeId" form:"animeId" query:"animeId"`
}
