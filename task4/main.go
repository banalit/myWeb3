package main

import (
	"github.com/gin-gonic/gin"
	task4Rest "github.com/luke/web3Learn/task4/rest"
)

func init() {

}

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	task4Rest.router = router
	task4Rest.AddRestApi(router)
	router.Run()
}
