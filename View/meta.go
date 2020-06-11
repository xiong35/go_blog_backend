package view

import (
	"go_blog/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
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
