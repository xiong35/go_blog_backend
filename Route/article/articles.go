package article

import (
	"go_blog/view"

	"github.com/gin-gonic/gin"
)

// HandleArticles dispatch requests for articles
func HandleArticles(r *gin.Engine) {
	artlcleGroup := r.Group("/articles")
	{
		artlcleGroup.GET("/", func(c *gin.Context) {
			c.String(200, "999")
		})
		artlcleGroup.POST("/tags", view.TagViewPost)
		artlcleGroup.POST("/trap", view.ArticleViewPost("trap"))
		artlcleGroup.POST("/blog", view.ArticleViewPost("blog"))
	}
}
