package handlers

import (
	"go_blog/database"
)

// ICan :form model for Can
type ICan struct {
	Icon  string `json:"icon" form:"icon"`
	Name  string `json:"name" form:"name"`
	Color string `json:"color" form:"color"`
}

// GetCan get all Can
func GetCan() (rtList []ICan) {

	database.DB.Table("i_cans").Find(&rtList)

	return
}

// Insert insert a Can
func (i_can *ICan) Insert() (err error) {

	result := database.DB.Table("i_cans").Create(&i_can)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
