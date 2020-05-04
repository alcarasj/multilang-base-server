package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const PORT = 8080

func index(c *gin.Context) {
	c.String(200, "Welcome! This is a base server written in Go using the Gin framework.")
}

func main() {
	port := strconv.Itoa(PORT)

	if port == "" {
		log.Fatal("Port must be set.")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", index)
	router.Run(":" + port)
}
