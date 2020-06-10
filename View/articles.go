package view

import (
	"go_blog/handlers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ArticleViewPost handle post requests for articles
func ArticleViewPost(t string) func(c *gin.Context) {

	return func(c *gin.Context) {
		var articleHandler handlers.Article
		var err error
		var id uint

		err = c.ShouldBindBodyWith(&articleHandler, binding.JSON)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err = articleHandler.Insert(t)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Insert() error!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"id":     id,
		})
	}

}

// TagViewPost handle post requests for articles
func TagViewPost(c *gin.Context) {
	var tagHandler handlers.Tag
	var err error
	var id uint

	err = c.ShouldBindBodyWith(&tagHandler, binding.JSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err = tagHandler.Insert()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Insert() error!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"id":     id,
	})

}

// ArticleViewGet handle get requests for articles
func ArticleViewGet(t string) func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.DefaultQuery("id", "0")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "unsupported id number",
			})
		}
		uid := uint(id)
		rtList := handlers.GetArticle(t, uid)

		if uid == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   rtList,
			})
			return
		}
		if len(rtList) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "id not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   rtList[0],
		})
		return
	}
}
