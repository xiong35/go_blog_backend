package article

import (
	"go_blog/middleware"
	"go_blog/view"

	"github.com/gin-gonic/gin"
)

// HandleArticles dispatch requests for articles
func HandleArticles(r *gin.Engine) {
	articleGroup := r.Group("/articles")
	articleGroup.Use(middleware.CheckAdmin())
	{
		// articleGroup.GET("/", func(c *gin.Context) {
		// 	c.String(200, "999")
		// })
		articleGroup.POST("/tags", view.TagViewPost)
		articleGroup.POST("/trap", view.ArticleViewPost("trap"))
		articleGroup.POST("/blog", view.ArticleViewPost("blog"))

		// articleGroup.GET("/tags", view.TagViewGet)
		articleGroup.GET("/trap", view.ArticleViewGet("trap"))
		articleGroup.GET("/blog", view.ArticleViewGet("blog"))
	}
}
