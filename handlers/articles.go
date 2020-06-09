package handlers

import (
	"go_blog/database"
	"go_blog/model"

	"github.com/jinzhu/gorm"
)

// Blog :form model for blog
type Blog struct {
	ID       uint     `json:"id" form:"id"`
	Content  string   `json:"content" form:"content"`
	Tags     []string `json:"tags" form:"tags[]"`
	Headline string   `json:"headline" form:"headline"`
}

// Insert :handler for blog
func (blog *Blog) Insert() (id uint, err error) {

	var blogModel model.Blog

	tags := database.DB.Table("tags").Where("tag_name in (?)", blog.Tags)
	tags.Find(&blogModel.Tags).Update("count", gorm.Expr("count + ?", 1))

	blogModel.Headline = blog.Headline
	blogModel.Content = blog.Content

	result := database.DB.Create(&blogModel)
	id = blogModel.ID
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
