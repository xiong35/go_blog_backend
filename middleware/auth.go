package middleware

import (
	"go_blog/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CheckAdmin :if no correct pw, post method is not allowed
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {

			var token handlers.Token
			err := c.ShouldBindBodyWith(&token, binding.JSON)

			rtToken, permission, new, err := token.Check()
			if new {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"status":     http.StatusUnauthorized,
					"token":      rtToken,
					"permission": permission,
				})
				return
			}

			fail := err != nil || permission < 9
			if fail {
				c.Abort()
				c.JSON(http.StatusOK, gin.H{
					"status":     http.StatusUnauthorized,
					"error":      "authorize failed",
					"token":      rtToken,
					"permission": permission,
				})
			}
		}
	}
}
