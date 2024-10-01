package admin_routes

import "github.com/gin-gonic/gin"

func AdminRoutes(router *gin.Engine) {
	adminGroup := router.Group("/admin")

	adminGroup.GET("/", func(c *gin.Context) {

	})
}
