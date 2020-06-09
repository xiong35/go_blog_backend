package database

// 连接数据库

import (
	"fmt"
	"go_blog/model"

	"github.com/jinzhu/gorm"
	// include mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 全局使用的数据库对象
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:xyl0321@/go_blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	DB.AutoMigrate(&model.Blog{}, &model.Tag{})
}
