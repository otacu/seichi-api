package main

import (
	"github.com/jinzhu/gorm"
)

type Anime struct {
	gorm.Model
	Name string
}

// 设置Anime的表名为`tb_anime`
func (Anime) TableName() string {
	return "tb_anime"
}

type Seichi struct {
	gorm.Model
	AnimeId uint
	Name    string
	Lng     float32
	Lat     float32
}

// 设置Seichi的表名为`tb_seichi`
func (Seichi) TableName() string {
	return "tb_seichi"
}
