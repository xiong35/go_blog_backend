package article

import (
	"go_blog/view"

	"github.com/gin-gonic/gin"
)

// HandleArticles dispatch requests for articles
func HandleArticles(r *gin.Engine) {
	artlcleGroup := r.Group("/articles")
	{
		artlcleGroup.POST("/tags", view.TagViewPost)
		artlcleGroup.POST("/blog", view.BlogViewPost)
	}
}
