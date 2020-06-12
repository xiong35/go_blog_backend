package view

import (
	"go_blog/database"
	"go_blog/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MetaViewGet handle get requests for meta
func MetaViewGet() func(c *gin.Context) {
	return func(c *gin.Context) {

		rtList := handlers.GetMeta()

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   rtList,
		})
		return
	}
}

// PostDuck handle duck kill
func PostDuck() func(c *gin.Context) {
	return func(c *gin.Context) {
		database.DB.Table("meta").Where("`key` = \"击杀小鸭\"").
			Update("value", gorm.Expr("value + ?", 1))
	}
}
