package view

import (
	"go_blog/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CanViewGet handle get requests for i cans
func CanViewGet() func(c *gin.Context) {
	return func(c *gin.Context) {

		rtList := handlers.GetCan()

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   rtList,
		})
		return
	}
}

// CanViewPost handle post requests for i cans
func CanViewPost() func(c *gin.Context) {
	return func(c *gin.Context) {

		var iCanHandler handlers.ICan

		err := c.ShouldBindBodyWith(&iCanHandler, binding.JSON)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = iCanHandler.Insert()

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Insert() error!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	}
}
