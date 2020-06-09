package route

import (
	"go_blog/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gin/Controllers"
	// "gin/Middlewares"
	// "gin/Sessions"
	// "github.com/gin-contrib/sessions"
)

// InitRouter init the router
func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middleware.Cors())
	// 使用 session(cookie-based)
	// router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("/pre")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "lololo")
		})
	}

	router.Run(":8080")
}
