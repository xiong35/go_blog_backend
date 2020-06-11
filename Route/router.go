package route

import (
	"go_blog/middleware"
	"go_blog/route/article"
	"go_blog/view"

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

	article.HandleArticles(router)

	router.GET("/meta", view.MetaViewGet())
	router.POST("/meta/duck", view.PostDuck())

	router.GET("/can", view.CanViewGet())
	router.POST("/can", view.CanViewPost())

	router.Run(":8080")
}
