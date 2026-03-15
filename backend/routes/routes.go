package routes

import (
	"backend/controllers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/health", healthHandler())
	r.HEAD("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := r.Group("/api")
	{
		api.GET("/members", controllers.GetMembers)
		api.POST("/members", controllers.CreateMember)
	}

	// r.POST("api/register", controllers.Register)
	// r.POST("api/login", controllers.Login)

	// protected := r.Group("/api")
	// protected.Use(middleware.AuthMiddleware())
	// {
	// 	protected.GET("/profile", controllers.Profile)
	// }
}

func healthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.1.0",
			"time":    time.Now().In(loc),
		})
	}
}
