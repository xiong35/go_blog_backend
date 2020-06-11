package middleware

import (
	"go_blog/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AddArticle add meta for each article
func AddArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 200 {
			database.DB.Table("meta").Where("`key` = \"article_num\"").
				Update("value", gorm.Expr("value + ?", 1))
		}
	}
}
