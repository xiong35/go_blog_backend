package model

import (
	"github.com/jinzhu/gorm"
)

// Tag model for tag
type Tag struct {
	ID      uint   `gorm:"primary_key"`
	TagName string `gorm:"type:char(7)"`
	Count   *uint
}

// Article model for trap
type Article struct {
	gorm.Model
	Content  string `gorm:"type:text"`
	Headline string `gorm:"type:varchar(200)"`
	Tags     []Tag  `gorm:"many2many:article_tag;"`
	Type     string
}
