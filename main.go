package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)

	router.Run(":8080")
}

func ping(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"application/json",
		[]byte(`{"message": "pong"}`),
	)
}
