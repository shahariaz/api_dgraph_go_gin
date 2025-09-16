package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	configDB "github.com/shahariaz/api_dgraph/config"
)

func main() {
	configDB.InitDgraph()
	r := gin.Default()

	// 3. Add routes
	// r.GET("/users/:id", handlers.GetUser)
	// r.POST("/users", handlers.CreateUser)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})
	// 4. Start server
	r.Run(":8081") // your app’s API port
}
