package model

import (
	"go_blog_system/constructor"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model //已有了创建时间和更新时间
	constructor.Article
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("testDB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&ArticleModel{})
}
