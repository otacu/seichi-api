package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

const (
	dataSource = "admin:123456@tcp(127.0.0.1:3306)/seichi?charset=utf8&parseTime=true"
)

// 数据库连接池
var DB *gorm.DB

func main() {
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Println(err)
		panic("连接数据库失败")
	}
	defer db.Close()
	DB = db

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 路径组 路径前缀
	r := e.Group("/seichi-api")
	r.GET("/getAllAnime", GetAllAnime)
	r.GET("/getSeichi", GetSeichi)

	// 启动
	e.Logger.Fatal(e.Start(":8896"))
}
