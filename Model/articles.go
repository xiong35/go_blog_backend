package model

import (
	"github.com/jinzhu/gorm"
)

// Blog model for blog
type Blog struct {
	gorm.Model
	Content  string `gorm:"type:text"`
	Headline string `gorm:"type:varchar(200)"`
	Tags     []Tag
	TagsID   uint
}

// Tag model for tag
type Tag struct {
	gorm.Model
	TagName string `gorm:"type:char(20)"`
	Count   *uint
}
