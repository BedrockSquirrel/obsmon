package main

import (
	"github.com/bedrocksquirrel/obsmon/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// REST Endpoints
	api := r.Group("/api")
	api.GET("/hello", handlers.HelloWorld)

	// Start server
	r.Run(":8080")
}
