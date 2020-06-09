package model

import (
	"github.com/jinzhu/gorm"
)

// Tag model for tag
type Tag struct {
	gorm.Model
	TagName string `gorm:"type:char(7)"`
	Count   *uint
}

// Article model for trap
type Article struct {
	gorm.Model
	Content  string `gorm:"type:text"`
	Headline string `gorm:"type:varchar(200)"`
	Tags     []Tag
	Type     string
}
