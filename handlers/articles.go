package handlers

import (
	"go_blog/database"
	"go_blog/model"

	"github.com/jinzhu/gorm"
)

// Article :form model for article
type Article struct {
	ID       uint     `json:"id" form:"id"`
	Content  string   `json:"content" form:"content"`
	Tags     []string `json:"tags" form:"tags[]"`
	Headline string   `json:"headline" form:"headline"`
}

// Insert :handler for article
func (article *Article) Insert(t string) (id uint, err error) {

	var articleModel model.Article

	tags := database.DB.Table("tags").Where("tag_name in (?)", article.Tags)
	tags.Find(&articleModel.Tags).Update("count", gorm.Expr("count + ?", 1))

	articleModel.Headline = article.Headline
	articleModel.Content = article.Content
	articleModel.Type = t

	result := database.DB.Create(&articleModel)
	id = articleModel.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Tag :form model for tag
type Tag struct {
	ID      uint   `json:"id" form:"id"`
	TagName string `json:"tag_name" form:"tag_name"`
}

// Insert :handler for tag
func (tag *Tag) Insert() (id uint, err error) {

	var tagModel model.Tag

	tagModel.TagName = tag.TagName
	tagModel.Count = new(uint)

	result := database.DB.Create(&tagModel)
	id = tagModel.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
