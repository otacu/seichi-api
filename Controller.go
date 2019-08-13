package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// 获取所有作品
func GetAllAnime(c echo.Context) error {
	var animeList []Anime
	DB.Find(&animeList)
	resultList := make([]GetAnimeResponse, len(animeList))
	for i := 0; i < len(animeList); i++ {
		resultList[i].Name = animeList[i].Name
		resultList[i].Id = animeList[i].ID
	}
	return c.JSON(http.StatusOK, resultList)
}

// 根据作品名称获取圣地
func GetSeichi(c echo.Context) error {
	animeId := c.QueryParam("animeId")
	var seichiList []Seichi
	DB.Find(&seichiList, "anime_id = ?", animeId)
	resultList := make([]GetSeichiResponse, len(seichiList))
	for i := 0; i < len(seichiList); i++ {
		resultList[i].Name = seichiList[i].Name
		resultList[i].Id = seichiList[i].ID
		resultList[i].AnimeId = seichiList[i].AnimeId
		resultList[i].Lng = seichiList[i].Lng
		resultList[i].Lat = seichiList[i].Lat
	}
	return c.JSON(http.StatusOK, resultList)
}
