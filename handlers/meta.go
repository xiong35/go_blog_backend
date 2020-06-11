package handlers

import (
	"go_blog/database"
)

// Meta :form model for meta
type Meta struct {
	Key      string `json:"key"`
	Value    uint   `json:"value"`
	Addition string `json:"addition "`
}

// GetMeta get all meta
func GetMeta() (rtList []Meta) {

	database.DB.Table("meta").Find(&rtList)

	return
}
