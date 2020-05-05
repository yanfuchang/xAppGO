package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()

	r.GET("index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"成功啦",
		})
	})

	r.Run(":9090")
	
}
