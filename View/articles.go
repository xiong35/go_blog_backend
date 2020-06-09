package view

import (
	"go_blog/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BlogViewPost handle post requests for articles
func BlogViewPost(c *gin.Context) {
	var blogHandler handlers.Blog
	var err error
	var id uint

	err = c.ShouldBind(&blogHandler)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err = blogHandler.Insert()
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

// TagViewPost handle post requests for articles
func TagViewPost(c *gin.Context) {
	var tagHandler handlers.Tag
	var err error
	var id uint

	err = c.ShouldBind(&tagHandler)

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
