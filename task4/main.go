package main

import (
	"log"
	middleware "task4/middleware"
	rest "task4/rest"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(LatencyLogger())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	AddRestApi(router)
	router.Run()
}

func AddRestApi(r *gin.Engine) {

	apiPublic := r.Group("/api")
	{
		apiPublic.POST("/register", rest.Register)
		apiPublic.POST("/login", rest.Login)
	}

	apiJwt := r.Group("/api")
	apiJwt.Use(middleware.JwtAuthMiddleware)
	{
		apiJwt.POST("/logout", rest.Logout)

		apiJwt.POST("/post", rest.CreatePost)
		apiJwt.DELETE("/post", rest.DeletePost)
		apiJwt.PATCH("/post", rest.PatchPost)
		apiJwt.GET("/post", rest.ListPost)
		apiJwt.POST("/comment", rest.CreateComment)
		apiJwt.DELETE("/comment", rest.DeleteComment)
		apiJwt.PATCH("/comment", rest.PatchComment)
		apiJwt.GET("/comment", rest.ListComment)
	}

}

func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		log.Printf("%s %s cost:%v", c.Request.Method, c.Request.URL.Path, latency)

	}
}
