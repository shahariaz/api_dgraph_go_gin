package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	configDB "github.com/shahariaz/api_dgraph/config"
)

func main() {
	configDB.InitDgraph()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})

	r.Run(":8081")
}
